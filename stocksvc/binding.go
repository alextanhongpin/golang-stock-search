package stocksvc

import (
	"encoding/json"
	"net/http"
	"strings"

	//httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)

type HTTPBinding struct {
	Ctx              context.Context
	GetStockEndpoint endpoint.Endpoint
}

func (b HTTPBinding) Register(r *httprouter.Router) {
	r.GET("/stocks", b.handleGet)
}

func (b HTTPBinding) handleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	queryValues := r.URL.Query()
	tickersSymbol := strings.Split(queryValues.Get("tickers_symbol"), ",")
	response, err := b.GetStockEndpoint(b.Ctx, getStockRequest{
		TickersSymbol: tickersSymbol,
	})
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondSuccess(w, response.(getStockResponse).Stocks)
}

// respondError in some canonical format.
func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":       err,
		"status_code": code,
		"status_text": http.StatusText(code),
	})
}
func respondSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
