package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

func quoteSingle(pairs []string, quotes map[string]GDAXTrade) {
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
		v.Trade.Price = round(v.Trade.Price)
		quotes[v.Name] = v.Trade
		if len(v.Trade.Price) > maxLength {
			maxLength = len(v.Trade.Price)
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
	print(spaces, quotes)
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
