package main

import (
	"flag"
)

func main() {
	stream := flag.Bool("s", false, "stream cryptocurrency prices")
	flag.Parse()

	state := make(map[string]Product, 9)

	pairs := []string{"BTC-USD", "BTC-EUR", "BTC-GBP", "ETH-USD", "ETH-BTC", "ETH-EUR", "LTC-USD", "LTC-BTC", "LTC-EUR"}

	for _, pair := range pairs {
		state[pair] = Product{ID: pair}
	}

	max := &MaxLengths{
		Price:  5,
		Size:   9,
		Bid:    3,
		Ask:    3,
		High:   4,
		Low:    3,
		Open:   4,
		Volume: 6,
	}

	if !*stream {
		quoteSingle(state, max)
	} else {
		// clear()
		// quoteStream(state, max, pairs)
	}
}
