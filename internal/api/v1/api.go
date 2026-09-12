package v1

import (
	"finance-tracker-backend/internal/api/v1/symbols"
	"net/http"
)

type Route struct{}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) RegisterHandlers(mux *http.ServeMux) {
	symbolsRoute := symbols.NewRoute()

	symbolsRoute.RegisterHandlers(mux)
}
