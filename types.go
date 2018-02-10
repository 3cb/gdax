package main

import "time"

// GDAXTrade contains data for a single trade
type GDAXTrade struct {
	Time    string `json:"time"`
	TradeID int64  `json:"trade_id"`
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
type Match struct {
	Type         string    `json:"type"`
	TradeID      int       `json:"trade_id"`
	Sequence     int       `json:"sequence"`
	MakerOrderID string    `json:"maker_order_id"`
	TakerOrderID string    `json:"taker_order_id"`
	Time         time.Time `json:"time"`
	ProductID    string    `json:"product_id"`
	Size         string    `json:"size"`
	Price        string    `json:"price"`
	Side         string    `json:"side"`
}
