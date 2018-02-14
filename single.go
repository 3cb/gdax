// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

func quoteSingle(state map[string]Product, max *MaxLengths) {
	tradeCh := make(chan Product, len(state))
	statsCh := make(chan Stats, len(state))
	tickerCh := make(chan Ticker, len(state))

	// concurrent http requests
	wg := &sync.WaitGroup{}
	wg.Add(3 * len(state))
	for _, pair := range state {
		go getTrades(pair.ID, tradeCh, wg)
		go getStats(pair.ID, statsCh, wg)
		go getTicker(pair.ID, tickerCh, wg)
	}
	wg.Wait()

	// set state from http response data
	for i := 0; i < len(state); i++ {
		product := Product{}

		trade := <-tradeCh
		product = state[trade.ID]
		product.Price = RndPrice(trade.Price)
		max.Price = getMax(max.Price, len(product.Price))
		product.Size = RndSize(trade.Size)
		max.Size = getMax(max.Size, len(product.Size))
		state[trade.ID] = product

		stats := <-statsCh
		product = state[stats.ID]
		product.Open = RndPrice(stats.Open)
		max.Open = getMax(max.Open, len(product.Open))
		product.High = RndPrice(stats.High)
		max.High = getMax(max.High, len(product.High))
		product.Low = RndPrice(stats.Low)
		max.Low = getMax(max.Low, len(product.Low))
		product.Volume = RndVol(stats.Volume)
		max.Volume = getMax(max.Volume, len(product.Volume))
		state[product.ID] = product

		ticker := <-tickerCh
		product = state[ticker.ID]
		product.Bid = RndPrice(ticker.Bid)
		max.Bid = getMax(max.Bid, len(product.Bid))
		product.Ask = RndPrice(ticker.Ask)
		max.Ask = getMax(max.Ask, len(product.Ask))
		state[product.ID] = product
	}

	// calculate max length of all deltas
	// set spacing on Delta
	for k, v := range state {
		v.Delta = SetDelta(v.Price, v.Open)
		v.Color = SetColor(v.Delta)
		max.Delta = getMax(max.Delta, len(v.Delta))
		state[k] = v
	}

	// format spacing
	for k, v := range state {
		v.Price = setSpc(max.Price, v.Price)
		v.Delta = setSpc(max.Delta, v.Delta)
		v.Size = setSpc(max.Size, v.Size)
		v.Bid = setSpc(max.Bid, v.Bid)
		v.Ask = setSpc(max.Ask, v.Ask)
		v.High = setSpc(max.High, v.High)
		v.Low = setSpc(max.Low, v.Low)
		v.Open = setSpc(max.Open, v.Open)
		v.Volume = setSpc(max.Volume, v.Volume)
		state[k] = v
	}

	clearScr()
	print(state, max)
}

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
