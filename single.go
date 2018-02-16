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
	for _, product := range state {
		go getTrades(product.ID, tradeCh, wg)
		go getStats(product.ID, statsCh, wg)
		go getTicker(product.ID, tickerCh, wg)
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
	for k, product := range state {
		product.Price = setSpc(max.Price, product.Price)
		product.Change = setSpc(max.Change, product.Change)
		product.Size = setSpc(max.Size, product.Size)
		product.Bid = setSpc(max.Bid, product.Bid)
		product.Ask = setSpc(max.Ask, product.Ask)
		product.High = setSpc(max.High, product.High)
		product.Low = setSpc(max.Low, product.Low)
		product.Open = setSpc(max.Open, product.Open)
		product.Volume = setSpc(max.Volume, product.Volume)
		product.Row = product.fmtRow()
		state[k] = product
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
