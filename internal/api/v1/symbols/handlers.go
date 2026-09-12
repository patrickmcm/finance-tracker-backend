package symbols

import (
	"fmt"
	"net/http"
)

func (route *Route) getSymbols(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello")
	if err != nil {
		return
	}
}

func (route *Route) getSymbolPriceData(w http.ResponseWriter, r *http.Request) {
	isin := r.PathValue("isin")
	// warning, unsanitised

	fmt.Fprintf(w, isin)
}
