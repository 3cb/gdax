// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

func quoteStream(state map[string]Product, pairs []string, max *MaxLengths) {
	format := quoteSingle(state, pairs, max)

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
			product.Price = setSpcStrm(max.Price, rndPrice(msg.Price))
			product.Size = setSpcStrm(max.Size, rndSize(msg.Size))
			product.Delta = setDelta(strings.TrimSpace(product.Price), strings.TrimSpace(product.Open))
			product.Color = SetColor(product.Delta)
			product.Delta = setSpcStrm(max.Delta, product.Delta)
			state[msg.ID] = product

		} else if msg.Type == "ticker" {
			product := state[msg.ID]
			product.Bid = setSpcStrm(max.Bid, rndPrice(msg.Bid))
			product.Ask = setSpcStrm(max.Ask, rndPrice(msg.Ask))
			product.High = setSpcStrm(max.High, rndPrice(msg.High))
			product.Low = setSpcStrm(max.Low, rndPrice(msg.Low))
			product.Open = setSpcStrm(max.Open, rndPrice(msg.Open))
			product.Volume = setSpcStrm(max.Volume, rndVol(msg.Volume))
			product.Delta = setDelta(strings.TrimSpace(product.Price), strings.TrimSpace(product.Open))
			product.Color = SetColor(product.Delta)
			product.Delta = setSpcStrm(max.Delta, product.Delta)
			state[msg.ID] = product
		}
		for k, v := range state {
			v.Row = fmtRow(v)
			state[k] = v
		}
		clearScr()
		print(state, format)
	}
}
