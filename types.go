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

// Match contains structure for websocket messages from matches channel
// Commented fields are available through API and remain for possible future use
type Match struct {
	Type      string `json:"type"`
	ProductID string `json:"product_id"`
	Price     string `json:"price"`
	Size      string `json:"size"`

	// TradeID   int    `json:"trade_id"`
	// Sequence     int    `json:"sequence"`
	// MakerOrderID string `json:"maker_order_id"`
	// TakerOrderID string `json:"taker_order_id"`
	// Time         string `json:"time"`
	// Side         string `json:"side"`
}

// Ticker contains structure for ticker websocket message
// Commented fields are available through API and remain for possible future use
// type Ticker struct {
// 	Type      string `json:"type"`
// 	ProductID string `json:"product_id"`
// 	BestBid   string `json:"best_bid"`
// 	BestAsk   string `json:"best_ask"`
// 	High24h   string `json:"high_24h"`
// 	Low24h    string `json:"low_24h"`
// 	Open24h   string `json:"open_24h"`
// 	Volume24h string `json:"volume_24h"`

// Volume30d string `json:"volume_30d"`
// Price     string `json:"price"`
// LastSize  string `json:"last_size"`
// Sequence  int64  `json:"sequence"`
// Side      string `json:"size"`
// Time      string `json:"time"`
// TradeID   int64  `json:"trade_id"`
// }

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

	// Fixed sets column width during quoteSingle so they don't change width while streaming from websocket
	Fixed int
}

func (m *MaxLengths) getTotal() int {
	return m.Price + m.Delta + m.Size + m.Bid + m.Ask + m.High + m.Low + m.Open + m.Volume
}
