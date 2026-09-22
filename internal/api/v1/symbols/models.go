package symbols

import "time"

type SymbolType string

const (
	ETF   SymbolType = "ETF"
	STOCK            = "STOCK"
)

type MarketSymbol struct {
	Ticker     string     `json:"ticker"`
	FullName   string     `json:"fullName"`
	Currency   string     `json:"currency"`
	Isin       string     `json:"isin"`
	SymbolType SymbolType `json:"symbolType"`
}

type MarketSymbolPriceData struct {
	timestamp time.Time
	marketAsk uint64
	marketBid uint64
}
