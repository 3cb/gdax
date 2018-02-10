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

		maxLength := 0
		for i := 0; i < 9; i++ {
			v := <-tradeCH
			quotes[v.Name] = v.Trade
			if len(v.Trade.Price) > maxLength {
				maxLength = len(v.Trade.Price)
			}
		}

		spaces := make(map[string]string)
		for k, v := range quotes {
			diff := maxLength - len(v.Price)
			s := "    "
			for i := 0; i < diff; i++ {
				s += " "
			}
			spaces[k] = s
		}
		println("\nGDAX Price Quotes:")
		fmt.Printf("\nBTC/USD:%v%v\n", spaces["BTC-USD"], quotes["BTC-USD"].Price)
		fmt.Printf("\nBTC/EUR:%v%v\n", spaces["BTC-EUR"], quotes["BTC-EUR"].Price)
		fmt.Printf("\nBTC/GBP:%v%v\n", spaces["BTC-GBP"], quotes["BTC-GBP"].Price)
		fmt.Printf("\nETH/USD:%v%v\n", spaces["ETH-USD"], quotes["ETH-USD"].Price)
		fmt.Printf("\nETH/BTC:%v%v\n", spaces["ETH-BTC"], quotes["ETH-BTC"].Price)
		fmt.Printf("\nETH/EUR:%v%v\n", spaces["ETH-EUR"], quotes["ETH-EUR"].Price)
		fmt.Printf("\nLTC/USD:%v%v\n", spaces["LTC-USD"], quotes["LTC-USD"].Price)
		fmt.Printf("\nLTC/BTC:%v%v\n", spaces["LTC-BTC"], quotes["LTC-BTC"].Price)
		fmt.Printf("\nLTC/EUR:%v%v\n", spaces["LTC-EUR"], quotes["LTC-EUR"].Price)
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
