package symbols

import "time"

type SymbolType int

const (
	ETF SymbolType = iota
	STOCK
)

type MarketSymbol struct {
	Isin       string     `json:"isin"`
	FullName   string     `json:"fullName"`
	Ticker     string     `json:"ticker"`
	SymbolType SymbolType `json:"symbolType"`
}

type MarketSymbolPriceData struct {
	timestamp time.Time
	marketAsk uint64
	marketBid uint64
}
