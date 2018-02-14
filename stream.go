// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

func quoteStream(state map[string]Product, max *MaxLengths, pairs []string) {
	columnHeaders := quoteSingle(state, max)

	wsSub := &Subscribe{
		Type:       "subscribe",
		ProductIds: pairs,
		Channels:   []string{"matches", "ticker"},
	}

	conn, resp, err := websocket.DefaultDialer.Dial("wss://ws-feed.gdax.com", nil)
	if resp.StatusCode != 101 || err != nil {
		log.Fatalf("Unable to connect to GDAX websocket API.")
	}
	conn.WriteJSON(wsSub)

	for {
		msg := Product{}

		err := conn.ReadJSON(&msg)
		if err != nil {
			conn.Close()
			clearScr()
			log.Fatalf("\nError from websocket: %s\nShutting down.", err)
		}

		if msg.Type == "match" {
			product := state[msg.ID]
			product.Price = SetSpcStrm(max.Price, RndPrice(msg.Price))
			product.Size = SetSpcStrm(max.Size, RndSize(msg.Size))
			product.Delta = SetDelta(strings.TrimSpace(product.Price), strings.TrimSpace(product.Open))
			product.Color = SetColor(product.Delta)
			product.Delta = SetSpcStrm(max.Delta, product.Delta)
			state[msg.ID] = product

		} else if msg.Type == "ticker" {
			product := state[msg.ID]
			product.Bid = SetSpcStrm(max.Bid, RndPrice(msg.Bid))
			product.Ask = SetSpcStrm(max.Ask, RndPrice(msg.Ask))
			product.High = SetSpcStrm(max.High, RndPrice(msg.High))
			product.Low = SetSpcStrm(max.Low, RndPrice(msg.Low))
			product.Open = SetSpcStrm(max.Open, RndPrice(msg.Open))
			product.Volume = SetSpcStrm(max.Volume, RndVol(msg.Volume))
			product.Delta = SetDelta(strings.TrimSpace(product.Price), strings.TrimSpace(product.Open))
			product.Color = SetColor(product.Delta)
			product.Delta = SetSpcStrm(max.Delta, product.Delta)
			state[msg.ID] = product
		}
		clearScr()
		print(state, columnHeaders)
	}
}
