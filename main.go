package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"github.com/alextanhongpin/stock-notifier/client"
	"github.com/alextanhongpin/stock-notifier/stocksvc"
	"golang.org/x/net/context"
)

// Yahoo: http://chartapi.finance.yahoo.com/instrument/1.0/msft/chartdata;type=quote;ys=2005;yz=4;ts=1234567890/json
// URL to get detailed company information for a single stock
// var urlDetailed string = "https://www.google.com/finance?q=JSE%3AIMP&q=JSE%3ANPN&ei=TrUBVomhAsKcUsP5mZAG&output=json"
// URL to get broad financials for multiple stocks
// https://www.quandl.com/api/v3/datasets/WIKI/FB.json
func main() {

	router := httprouter.New()
	ctx := context.Background()

	// Register the stock service
	svc := stocksvc.NewStockService()
	stocksvc.HTTPBinding{
		Ctx:              ctx,
		GetStockEndpoint: stocksvc.MakeGetStockEndpoint(svc),
	}.Register(router)

	client.RegisterRoutes(router)

	// return
	log.Fatal(http.ListenAndServe(":8080", router))
}
