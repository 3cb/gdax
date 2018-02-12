package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

func quoteSingle(state map[string]Product, max *MaxLengths) {
	tradeCh := make(chan Product, 9)
	statsCh := make(chan Stats, 9)
	tickerCh := make(chan Ticker, 9)

	wg := &sync.WaitGroup{}
	wg.Add(27)
	for _, pair := range state {
		go getTrades(pair.ID, tradeCh, wg)
		go getStats(pair.ID, statsCh, wg)
		go getTicker(pair.ID, tickerCh, wg)
	}
	wg.Wait()

	// set state from http response data
	for i := 0; i < 9; i++ {
		product := Product{}

		trade := <-tradeCh
		product = state[trade.ID]
		product.Price = rndPrice(trade.Price)
		max.Price = getMax(max.Price, len(product.Price))
		product.Size = rndSize(trade.Size)
		max.Size = getMax(max.Size, len(product.Size))
		state[trade.ID] = product

		stats := <-statsCh
		product = state[stats.ID]
		product.Open = rndPrice(stats.Open)
		max.Open = getMax(max.Open, len(product.Open))
		product.High = rndPrice(stats.High)
		max.High = getMax(max.High, len(product.High))
		product.Low = rndPrice(stats.Low)
		max.Low = getMax(max.Low, len(product.Low))
		product.Volume = rndVol(stats.Volume)
		max.Volume = getMax(max.Volume, len(product.Volume))
		state[product.ID] = product

		ticker := <-tickerCh
		product = state[ticker.ID]
		product.Bid = rndPrice(ticker.Bid)
		max.Bid = getMax(max.Bid, len(product.Bid))
		product.Ask = rndPrice(ticker.Ask)
		max.Ask = getMax(max.Ask, len(product.Ask))
		state[product.ID] = product
	}

	// calculate max length of all deltas
	// set spacing on Delta
	for k, v := range state {
		v.Delta, v.Color = setDeltaColor(v.Price, v.Open)
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
