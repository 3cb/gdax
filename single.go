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
		state[trade.ID] = processTradeCh(state, max, &trade)

		stats := <-statsCh
		state[stats.ID] = processStatsCh(state, max, &stats)

		ticker := <-tickerCh
		state[ticker.ID] = processTickerCh(state, max, &ticker)
	}

	// calculate max length of all deltas
	// set spacing on Delta
	for k, v := range state {
		v.Delta = SetDelta(v.Price, v.Open)
		v.Color = SetColor(v.Delta)
		max.Delta = SetMax(max.Delta, len(v.Delta))
		state[k] = v
	}

	// format spacing
	for k, v := range state {
		v.Price = SetSpc(max.Price, v.Price)
		v.Delta = SetSpc(max.Delta, v.Delta)
		v.Size = SetSpc(max.Size, v.Size)
		v.Bid = SetSpc(max.Bid, v.Bid)
		v.Ask = SetSpc(max.Ask, v.Ask)
		v.High = SetSpc(max.High, v.High)
		v.Low = SetSpc(max.Low, v.Low)
		v.Open = SetSpc(max.Open, v.Open)
		v.Volume = SetSpc(max.Volume, v.Volume)
		v.Row = FmtRow(v)
		state[k] = v
	}

	format := &FmtPrint{}
	format.Headers = FmtColHdr(max)
	format.Title = FmtTitle("GDAX Cryptocurrency Exchange", len(format.Headers))
	format.Footer = FmtTitle("", len(format.Headers))
	clearScr()
	print(state, format)
	return format
}

func processTradeCh(state map[string]Product, max *MaxLengths, trade *Product) Product {
	product := state[trade.ID]
	product.Price = RndPrice(trade.Price)
	max.Price = SetMax(max.Price, len(product.Price))
	product.Size = RndSize(trade.Size)
	max.Size = SetMax(max.Size, len(product.Size))
	return product
}

func processStatsCh(state map[string]Product, max *MaxLengths, stats *Stats) Product {
	product := state[stats.ID]
	product.Open = RndPrice(stats.Open)
	max.Open = SetMax(max.Open, len(product.Open))
	product.High = RndPrice(stats.High)
	max.High = SetMax(max.High, len(product.High))
	product.Low = RndPrice(stats.Low)
	max.Low = SetMax(max.Low, len(product.Low))
	product.Volume = RndVol(stats.Volume)
	max.Volume = SetMax(max.Volume, len(product.Volume))
	return product
}

func processTickerCh(state map[string]Product, max *MaxLengths, ticker *Ticker) Product {
	product := state[ticker.ID]
	product.Bid = RndPrice(ticker.Bid)
	max.Bid = SetMax(max.Bid, len(product.Bid))
	product.Ask = RndPrice(ticker.Ask)
	max.Ask = SetMax(max.Ask, len(product.Ask))
	return product
}
