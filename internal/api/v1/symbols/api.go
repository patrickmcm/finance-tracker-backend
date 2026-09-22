package symbols

import (
	"database/sql"
	"net/http"
)

type Route struct {
	db *sql.DB
}

func NewRoute(db *sql.DB) *Route {
	return &Route{db: db}
}

func (route *Route) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/symbol/", route.getSymbol)
	mux.HandleFunc("GET /api/v1/symbol/{ticker}", route.getSymbol)
	mux.HandleFunc("GET /api/v1/symbol/{ticker}/pricedata", route.getSymbolPriceData)
}
