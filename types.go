package main

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
