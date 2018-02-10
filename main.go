package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
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

func getTrades(pair string, tradeCH chan<- Currency, wg *sync.WaitGroup) {
	slice := []GDAXTrade{}

	api := "https://api.gdax.com/products/" + pair + "/trades?limit=1"
	resp, err := http.Get(api)
	if err != nil {
		tradeCH <- Currency{pair, "-----"}
		wg.Done()
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tradeCH <- Currency{pair, "-----"}
		wg.Done()
	}
	err = json.Unmarshal(bytes, &slice)
	if err != nil {
		tradeCH <- Currency{pair, "-----"}
		wg.Done()
	}
	tradeCH <- Currency{pair, slice[0].Price}
	wg.Done()
}
