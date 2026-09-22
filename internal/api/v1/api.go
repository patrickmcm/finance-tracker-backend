package v1

import (
	"database/sql"
	"finance-tracker-backend/internal/api/v1/symbols"
	"net/http"
)

type Route struct {
	db *sql.DB
}

func NewRoute(db *sql.DB) *Route {
	return &Route{db: db}
}

func (r *Route) RegisterHandlers(mux *http.ServeMux) {
	symbolsRoute := symbols.NewRoute(r.db)

	symbolsRoute.RegisterHandlers(mux)
}
