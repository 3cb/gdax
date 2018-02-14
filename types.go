// Command gdax gets a single quote from http requests to https://api.gdax.com or streams quotes from websocket at wss://ws-feed.gdax.com
package main

import "github.com/fatih/color"

// Product maintains up to date price and volume data for a cryptocurrency pair
// Trailing comments denote which http request or websocket stream the data comes from
// getTrades: https://docs.gdax.com/#get-trades
// match: https://docs.gdax.com/#the-code-classprettyprintfullcode-channel
// ticker: https://docs.gdax.com/#the-code-classprettyprinttickercode-channel
// *** GDAX API documentation for websocket ticker channel does not show all available fields as of 2/11/2018
type Product struct {
	Type string `json:"type"`

	ID    string       `json:"product_id"`
	Price string       `json:"price"` // getTrades/match
	Delta string       // % change in price
	Color *color.Color // set along with delta to color data on display
	Size  string       `json:"size"` // getTrades/match

	Bid    string `json:"best_bid"`   // ticker
	Ask    string `json:"best_ask"`   // ticker
	High   string `json:"high_24h"`   // ticker
	Low    string `json:"low_24h"`    // ticker
	Open   string `json:"open_24h"`   // ticker
	Volume string `json:"volume_24h"` // ticker

	Row string
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

// Subscribe is the structure for the subscription message sent to GDAX websocket API
// https://docs.gdax.com/#subscribe
type Subscribe struct {
	Type       string   `json:"type"`
	ProductIds []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

// MaxLengths is used to format spaces for printing
type MaxLengths struct {
	Price  int
	Delta  int
	Size   int
	Bid    int
	Ask    int
	High   int
	Low    int
	Open   int
	Volume int
}

// FmtPrint contains structure for printing so it only have to be calculated once
type FmtPrint struct {
	Title   string
	Headers string
	Footer  string
}
