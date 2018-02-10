package main

// GDAXTrade contains data for a single trade
type GDAXTrade struct {
	Time    string `json:"time"`
	TradeID int64  `json:"trade_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Side    string `json:"side"`
}

// Currency contains name and price of crypto currency only
type Currency struct {
	Name  string
	Price string
}
