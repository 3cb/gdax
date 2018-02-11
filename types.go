package main

// State maintains up to date price and volume data for each cryptocurrency
type State struct {
	ProductID   string
	DisplayName string
	Price       string
	BestBid     string
	BestAsk     string
	High24h     string
	Low24h      string
	Open24h     string
	Volume24h   string
}

// GDAXTrade contains data for a single trade
type GDAXTrade struct {
	Time    string `json:"time"`
	TradeID int    `json:"trade_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Side    string `json:"side"`
}

// Currency contains name and trade data of cryptocurrency
type Currency struct {
	Name  string
	Trade GDAXTrade
}

// Subscribe is the structure for the subscription message sent to GDAX websocket API
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

	// Size      string `json:"size"`
	// TradeID   int    `json:"trade_id"`
	// Sequence     int    `json:"sequence"`
	// MakerOrderID string `json:"maker_order_id"`
	// TakerOrderID string `json:"taker_order_id"`
	// Time         string `json:"time"`
	// Side         string `json:"side"`
}

// Ticker contains structure for ticker websocket message
// Commented fields are available through API and remain for possible future use
type Ticker struct {
	Type      string `json:"type"`
	ProductID string `json:"product_id"`
	BestBid   string `json:"best_bid"`
	BestAsk   string `json:"best_ask"`
	High24h   string `json:"high_24h"`
	Low24h    string `json:"low_24h"`
	Open24h   string `json:"open_24h"`
	Volume24h string `json:"volume_24h"`

	// Volume30d string `json:"volume_30d"`
	// Price     string `json:"price"`
	// LastSize  string `json:"last_size"`
	// Sequence  int64  `json:"sequence"`
	// Side      string `json:"size"`
	// Time      string `json:"time"`
	// TradeID   int64  `json:"trade_id"`
}
