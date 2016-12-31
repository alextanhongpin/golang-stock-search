package stocksvc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type getStockRequest struct {
	TickersSymbol []string `json:"tickers_symbol`
}

type getStockResponse struct {
	Stocks []Stock `json:"stocks"`
	Err    string  `json:"error,omitempty"`
}

func MakeGetStockEndpoint(svc StockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getStockRequest)
		var stocks []Stock
		//var resp []StockResponse
		symbolString := strings.Join(req.TickersSymbol, ",")
		var urlStocks string = "https://www.google.com/finance/info?infotype=infoquoteall&q=" + symbolString

		res, err := http.Get(urlStocks)

		if err != nil {
			return getStockResponse{stocks, err.Error()}, nil
		}
		// v, err :=

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getStockResponse{stocks, err.Error()}, nil
		}
		body = sanitizeBody(body)
		err = json.Unmarshal(body, &stocks)
		if err != nil {
			return getStockResponse{stocks, err.Error()}, nil
		}
		//StockResponse = stocks
		fmt.Println(stocks)
		return getStockResponse{stocks, ""}, nil
	}
}

func sanitizeBody(body []byte) (responseBody []byte) {

	body = body[4 : len(body)-1]
	body = bytes.Replace(body, []byte("\\x2F"), []byte("/"), -1)
	body = bytes.Replace(body, []byte("\\x26"), []byte("&"), -1)
	body = bytes.Replace(body, []byte("\\x3B"), []byte(";"), -1)
	body = bytes.Replace(body, []byte("\\x27"), []byte("'"), -1)

	responseBody = body
	return responseBody
}
