// // Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

func getTrades(pair string, tradeCh chan<- Product, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	slice := []Product{}

	api := "https://api.gdax.com/products/" + pair + "/trades?limit=1"
	resp, err := http.Get(api)
	if err != nil {
		tradeCh <- Product{ID: pair, Price: "-----", Size: "-----"}
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tradeCh <- Product{ID: pair, Price: "-----", Size: "-----"}
		return
	}
	err = json.Unmarshal(bytes, &slice)
	if err != nil {
		tradeCh <- Product{ID: pair, Price: "-----", Size: "-----"}
		return
	}
	slice[0].ID = pair
	tradeCh <- slice[0]
}

func getStats(pair string, statsCh chan<- Stats, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	stats := Stats{}

	api := "https://api.gdax.com/products/" + pair + "/stats"
	resp, err := http.Get(api)
	if err != nil {
		statsCh <- Stats{ID: pair, Open: "-----", High: "-----", Low: "-----", Volume: "-----"}
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		statsCh <- Stats{ID: pair, Open: "-----", High: "-----", Low: "-----", Volume: "-----"}
		return
	}
	err = json.Unmarshal(bytes, &stats)
	if err != nil {
		statsCh <- Stats{ID: pair, Open: "-----", High: "-----", Low: "-----", Volume: "-----"}
		return
	}
	stats.ID = pair
	statsCh <- stats
}

func getTicker(pair string, tickerCh chan<- Ticker, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	ticker := Ticker{}

	api := "https://api.gdax.com/products/" + pair + "/ticker"
	resp, err := http.Get(api)
	if err != nil {
		tickerCh <- Ticker{ID: pair, Bid: "-----", Ask: "-----"}
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tickerCh <- Ticker{ID: pair, Bid: "-----", Ask: "-----"}
		return
	}
	err = json.Unmarshal(bytes, &ticker)
	if err != nil {
		tickerCh <- Ticker{ID: pair, Bid: "-----", Ask: "-----"}
		return
	}
	ticker.ID = pair
	tickerCh <- ticker
}
