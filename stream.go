package main

import (
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

func quoteStream(state map[string]Product, max *MaxLengths, pairs []string) {
	quoteSingle(state, max)

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
			product.Price = setSpc(max.Price, rndPrice(msg.Price))
			product.Size = setSpc(max.Size, rndSize(msg.Size))
			product.Delta, product.Color = setDeltaColor(strings.TrimSpace(product.Price), strings.TrimSpace(product.Open))
			product.Delta = setSpc(max.Delta, product.Delta)
			state[msg.ID] = product

		} else if msg.Type == "ticker" {
			product := state[msg.ID]
			product.Bid = setSpc(max.Bid, rndPrice(msg.Bid))
			product.Ask = setSpc(max.Ask, rndPrice(msg.Ask))
			product.High = setSpc(max.High, rndPrice(msg.High))
			product.Low = setSpc(max.Low, rndPrice(msg.Low))
			product.Open = setSpc(max.Open, rndPrice(msg.Open))
			product.Volume = setSpc(max.Volume, rndVol(msg.Volume))
			product.Delta, product.Color = setDeltaColor(strings.TrimSpace(product.Price), strings.TrimSpace(product.Open))
			product.Delta = setSpc(max.Delta, product.Delta)
			state[msg.ID] = product
		}
		clearScr()
		print(state, max)
	}
}
