// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import (
	"strings"

	"github.com/fatih/color"
)

// Product maintains up to date price and volume data for a cryptocurrency pair
// Trailing comments denote which http request or websocket stream the data comes from
// getTrades: https://docs.gdax.com/#get-trades
// match: https://docs.gdax.com/#the-code-classprettyprintfullcode-channel
// ticker: https://docs.gdax.com/#the-code-classprettyprinttickercode-channel
// *** GDAX API documentation for websocket ticker channel does not show all available fields as of 2/11/2018
type Product struct {
	Type string `json:"type"`

	ID    string `json:"product_id"`
	Price string `json:"price"` // getTrades/match
	Size  string `json:"size"`  // getTrades/match

	High   string `json:"high_24h"`   // getStats/ticker
	Low    string `json:"low_24h"`    // getStats/ticker
	Open   string `json:"open_24h"`   // getStats/ticker
	Volume string `json:"volume_24h"` // getStats/ticker

	Bid string `json:"best_bid"` // getTicker/ticker
	Ask string `json:"best_ask"` // getTicker/ticker

	Change string       // % change in price
	Color  *color.Color // set along with delta to color data on display

	Row string
}

// fmtRow formats all Product data into a single row so it can be printed
func (p Product) fmtRow() string {
	b := strings.Builder{}
	b.WriteString(" ")
	b.WriteString(strings.Join(strings.Split(p.ID, "-"), "/"))
	b.WriteString(p.Price)
	b.WriteString(p.Size)
	b.WriteString(p.Change)
	b.WriteString(p.Bid)
	b.WriteString(p.Ask)
	b.WriteString(p.High)
	b.WriteString(p.Low)
	b.WriteString(p.Volume)
	return b.String()
}

// Stats contains 24 hour data from REST API:
// https://docs.gdax.com/#get-24hr-stats
type Stats struct {
	ID     string
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Volume string `json:"volume"`
}

// Ticker contains snapshot data from REST API:
// https://docs.gdax.com/#get-product-ticker
type Ticker struct {
	ID  string
	Bid string `json:"bid"`
	Ask string `json:"ask"`
}

// setSpc adds space in front of string based on max length for single quote
func setSpc(max int, orig string) string {
	b := strings.Builder{}
	b.WriteString("     ")
	diff := max - len(orig)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			b.WriteString(" ")
		}
	}
	b.WriteString(orig)
	return b.String()
}

// setSpcStrm adds spaces to string according to max lengths calculated in single quote
func setSpcStrm(max int, orig string) string {
	b := strings.Builder{}
	diff := 5 + max - len(orig)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			b.WriteString(" ")
		}
	}
	b.WriteString(orig)
	return b.String()
}
