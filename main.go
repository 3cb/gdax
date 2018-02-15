// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"flag"
)

func main() {
	stream := flag.Bool("s", false, "stream cryptocurrency prices")
	flag.Parse()

	pairs := []string{"BTC-USD", "BTC-EUR", "BTC-GBP", "BCH-USD", "BCH-BTC", "BCH-EUR", "ETH-USD", "ETH-BTC", "ETH-EUR", "LTC-USD", "LTC-BTC", "LTC-EUR"}

	state := make(map[string]Product, len(pairs))

	for _, pair := range pairs {
		state[pair] = Product{ID: pair}
	}

	// initialize with header lengths
	max := &MaxLengths{
		Price:  5,
		Change: 6,
		Size:   9,
		Bid:    3,
		Ask:    3,
		High:   4,
		Low:    3,
		Open:   4,
		Volume: 6,
	}

	if !*stream {
		quoteSingle(state, pairs, max)
	} else {
		quoteStream(state, pairs, max)
	}
}
