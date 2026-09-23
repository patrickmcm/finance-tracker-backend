package symbols

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"finance-tracker-backend/internal/api/v1/util"
	"fmt"
	"net/http"
)

func (route *Route) getInstrument(w http.ResponseWriter, r *http.Request) {
	ticker := r.PathValue("ticker")
	if ticker == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := route.db.Query("SELECT * FROM instruments WHERE ticker =$1", ticker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	instruments := util.GetTable[MarketSymbol](rows)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(instruments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (route *Route) getInstruments(w http.ResponseWriter, r *http.Request) {
	rows, err := route.db.Query("SELECT * FROM instruments")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	instruments := util.GetTable[MarketSymbol](rows)

	out, err := json.Marshal(instruments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hash := sha256.Sum256(out)

	eTag := fmt.Sprintf(`"%s"`, base64.RawURLEncoding.EncodeToString(hash[:]))

	if r.Header.Get("If-None-Match") == eTag {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("ETag", eTag)
	w.Write(out)
}

func (route *Route) getSymbolPriceData(w http.ResponseWriter, r *http.Request) {
	isin := r.PathValue("ticker")

	fmt.Fprintf(w, isin)
}
