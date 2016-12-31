package stocksvc

type Stock struct {
	Symbol           string `json:"t"`
	Exchange         string `json:"e"`
	Name             string `json:"name"`
	Change           string `json:"c"`
	Close            string `json:"l"`
	PercentageChange string `json:"cp"`
	Open             string `json:"op"`
	High             string `json:"hi"`
	Low              string `json:"lo"`
	Volume           string `json:"vo"`
	AverageVolume    string `json:"avvo"`
	High52           string `json:"hi52"`
	Low52            string `json:"lo52"`
	MarketCap        string `json:"mc"`
	EPS              string `json:"eps"`
	Shares           string `json:"shares"`
}

type StockService interface {
	GetStocks([]string) (stocks []Stock)
}

func NewStockService() StockService {
	return stockService{}
}

type stockService struct{}

func (svc stockService) GetStocks(stocks []string) (stockResponse []Stock) {
	return stockResponse
}
