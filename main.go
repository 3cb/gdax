package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	stream := flag.Bool("s", false, "stream cryptocurrency prices")
	flag.Parse()

	pairs := []string{"BTC-USD", "BTC-EUR", "BTC-GBP", "ETH-USD", "ETH-BTC", "ETH-EUR", "LTC-USD", "LTC-BTC", "LTC-EUR"}

	quotes := make(map[string]GDAXTrade)

	if !*stream {
		tradeCH := make(chan Currency, 9)
		wg := &sync.WaitGroup{}
		wg.Add(9)
		for _, pair := range pairs {
			go getTrades(pair, tradeCH, wg)
		}
		wg.Wait()

		for i := 0; i < 9; i++ {
			v := <-tradeCH
			quotes[v.Name] = v.Trade
		}
		println("\nGDAX Price Quotes:")
		fmt.Printf("\nBTC/USD: %v\n", quotes["BTC-USD"].Price)
		fmt.Printf("\nBTC/EUR: %v\n", quotes["BTC-EUR"].Price)
		fmt.Printf("\nBTC/GBP: %v\n", quotes["BTC-GBP"].Price)
		fmt.Printf("\nETH/USD: %v\n", quotes["ETH-USD"].Price)
		fmt.Printf("\nETH/BTC: %v\n", quotes["ETH-BTC"].Price)
		fmt.Printf("\nETH/EUR: %v\n", quotes["ETH-EUR"].Price)
		fmt.Printf("\nLTC/USD: %v\n", quotes["LTC-USD"].Price)
		fmt.Printf("\nLTC/BTC: %v\n", quotes["LTC-BTC"].Price)
		fmt.Printf("\nLTC/EUR: %v\n", quotes["LTC-EUR"].Price)
	}

}

func getTrades(pair string, tradeCH chan<- Currency, wg *sync.WaitGroup) {
	slice := []GDAXTrade{}

	api := "https://api.gdax.com/products/" + pair + "/trades?limit=1"
	resp, err := http.Get(api)
	if err != nil {
		tradeCH <- Currency{Name: pair}
		wg.Done()
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tradeCH <- Currency{Name: pair}
		wg.Done()
	}
	err = json.Unmarshal(bytes, &slice)
	if err != nil {
		tradeCH <- Currency{Name: pair}
		wg.Done()
	}
	tradeCH <- Currency{pair, slice[0]}
	wg.Done()
}
