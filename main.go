package main

import (
	"flag"
	"net/http"
	"sync"
)

func main() {
	stream := flag.Bool("s", false, "stream cryptocurrency prices")
	flag.Parse()

	pairs := []string{"BTC-USD", "BTC-EUR", "BTC-GBP", "ETH-USD", "ETH-BTC", "ETH-EUR", "LTC-USD", "LTC-BTC", "LTC-EUR"}

	tradeCH := make(chan Currency, 9)

	if !*stream {
		wg := &sync.WaitGroup{}
		wg.Add(9)
		for _, pair := range pairs {
			go getTrades(pair, tradeCH, wg)
		}
		wg.Wait()
	}

}

func getTrades(pair string, trade chan<- Currency, wg *sync.WaitGroup) {
	api := "https://api.gdax.com/products/" + pair + "/trades?limit=1"
	resp, err := http.Get(api)
	if err != nil {

	}
	wg.Done()
}
