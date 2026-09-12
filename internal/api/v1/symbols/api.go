package symbols

import "net/http"

type Route struct{}

func NewRoute() *Route {
	return &Route{}
}

func (route *Route) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/symbols", route.getSymbols)
	mux.HandleFunc("GET /api/v1/symbols/{isin}/pricedata", route.getSymbolPriceData)
}
