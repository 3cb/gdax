package main

func quoteStream(state map[string]Product, max *MaxLengths, pairs []string) {
	// quoteSingle(state, max)

	// wsSub := &Subscribe{
	// 	Type:       "subscribe",
	// 	ProductIds: pairs,
	// 	Channels:   []string{"matches", "ticker"},
	// }

	// conn, resp, err := websocket.DefaultDialer.Dial("wss://ws-feed.gdax.com", nil)
	// if resp.StatusCode != 101 || err != nil {
	// 	log.Fatalf("Unable to connect to GDAX websocket API.")
	// }
	// conn.WriteJSON(wsSub)

	// msg := Match{}
	// trade := GDAXTrade{}

	// for {
	// 	err := conn.ReadJSON(&msg)
	// 	if err != nil {
	// 		conn.Close()
	// 		log.Fatalf("Error from websocket: %s\nShutting down.", err)
	// 	}

	// 	if msg.Type == "match" {
	// 		trade.Price = round(msg.Price)
	// 		// trade.TradeID = msg.TradeID
	// 		// trade.Size = msg.Size
	// 		// trade.Time = msg.Time
	// 		// trade.Side = msg.Side

	// 		maxLength := 0
	// 		quotes[msg.ProductID] = trade
	// 		for _, v := range quotes {
	// 			if len(v.Price) > maxLength {
	// 				maxLength = len(v.Price)
	// 			}
	// 		}

	// 		spaces := make(map[string]string)
	// 		for k, v := range quotes {
	// 			diff := maxLength - len(v.Price)
	// 			s := "    "
	// 			if diff > 0 {
	// 				for i := 0; i < diff; i++ {
	// 					s += " "
	// 				}
	// 			}
	// 			spaces[k] = s
	// 		}
	// 		clear()
	// 		print(spaces, quotes)
	// 	} else if msg.Type == "ticker" {

	// 	}
	// }
}
