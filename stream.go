package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func quoteStream(pairs []string) {
	quotes := make(map[string]GDAXTrade)
	quoteSingle(pairs, quotes)

	wsSub := &Subscribe{
		Type:       "subscribe",
		ProductIds: pairs,
		Channels:   []string{"matches"},
	}

	conn, resp, err := websocket.DefaultDialer.Dial("wss://ws-feed.gdax.com", nil)
	if resp.StatusCode != 101 || err != nil {
		log.Fatalf("Unable to connect to GDAX websocket API.")
	}
	conn.WriteJSON(wsSub)

	msg := Match{}
	trade := GDAXTrade{}

	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			conn.Close()
			log.Fatalf("Error from websocket: %s\nShutting down.", err)
		}

		if msg.Type == "match" {
			trade.Time = msg.Time
			trade.TradeID = msg.TradeID
			trade.Price = msg.Price
			trade.Size = msg.Size
			trade.Side = msg.Side

			maxLength := 0
			quotes[msg.ProductID] = trade
			for _, v := range quotes {
				if len(v.Price) > maxLength {
					maxLength = len(v.Price)
				}
			}

			spaces := make(map[string]string)
			for k, v := range quotes {
				diff := maxLength - len(v.Price)
				s := "    "
				if diff > 0 {
					for i := 0; i < diff; i++ {
						s += " "
					}
				}
				spaces[k] = s
			}
			clear()
			print(spaces, quotes)
		}
	}
}
