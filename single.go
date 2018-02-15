// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"sync"
)

func quoteSingle(state map[string]Product, pairs []string, max *MaxLengths) *FmtPrint {
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
		trade := <-tradeCh
		state[trade.ID] = processTradeCh(state, &trade)

		stats := <-statsCh
		state[stats.ID] = processStatsCh(state, &stats)

		ticker := <-tickerCh
		state[ticker.ID] = processTickerCh(state, &ticker)
	}

	// calculate max lengths
	for _, product := range state {
		max.Price = setMax(max.Price, len(product.Price))
		max.Size = setMax(max.Size, len(product.Size))
		max.Open = setMax(max.Open, len(product.Open))
		max.High = setMax(max.High, len(product.High))
		max.Low = setMax(max.Low, len(product.Low))
		max.Volume = setMax(max.Volume, len(product.Volume))
		max.Bid = setMax(max.Bid, len(product.Bid))
		max.Ask = setMax(max.Ask, len(product.Ask))
		max.Change = setMax(max.Change, len(product.Change))
	}

	// format spacing
	for k, v := range state {
		v.Price = setSpc(max.Price, v.Price)
		v.Change = setSpc(max.Change, v.Change)
		v.Size = setSpc(max.Size, v.Size)
		v.Bid = setSpc(max.Bid, v.Bid)
		v.Ask = setSpc(max.Ask, v.Ask)
		v.High = setSpc(max.High, v.High)
		v.Low = setSpc(max.Low, v.Low)
		v.Open = setSpc(max.Open, v.Open)
		v.Volume = setSpc(max.Volume, v.Volume)
		v.Row = fmtRow(v)
		state[k] = v
	}

	format := &FmtPrint{}
	format.Headers = fmtColHdr(max)
	format.Title = fmtTitle("GDAX Cryptocurrency Exchange", len(format.Headers))
	format.Footer = fmtTitle("", len(format.Headers))
	clearScr()
	print(state, format)
	return format
}

func processTradeCh(state map[string]Product, trade *Product) Product {
	product := state[trade.ID]
	product.Price = rndPrice(trade.Price)
	product.Size = rndSize(trade.Size)
	if len(product.Open) > 0 {
		product.Change = setDelta(product.Price, product.Open)
		product.Color = SetColor(product.Change)
	}
	return product
}

func processStatsCh(state map[string]Product, stats *Stats) Product {
	product := state[stats.ID]
	product.Open = rndPrice(stats.Open)
	product.High = rndPrice(stats.High)
	product.Low = rndPrice(stats.Low)
	product.Volume = rndVol(stats.Volume)
	if len(product.Price) > 0 {
		product.Change = setDelta(product.Price, product.Open)
		product.Color = SetColor(product.Change)
	}
	return product
}

func processTickerCh(state map[string]Product, ticker *Ticker) Product {
	product := state[ticker.ID]
	product.Bid = rndPrice(ticker.Bid)
	product.Ask = rndPrice(ticker.Ask)
	return product
}
