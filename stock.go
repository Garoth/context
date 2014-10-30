package main

// https://code.google.com/p/yahoo-finance-managed/wiki/enumQuoteProperty
// https://code.google.com/p/yahoo-finance-managed/wiki/csvQuotesDownload

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type StockInfo struct {
	Symbol               string
	Name                 string
	LastTradePrice       float64
	OpeningPrice         float64
	PreviousClosingPrice float64
}

func NewStockInfo(symbol string) (*StockInfo, error) {
	var urlTmpl = "http://download.finance.yahoo.com/d/quotes.csv?s=%s&f=nl1op"
	url := fmt.Sprintf(urlTmpl, symbol)

	resp, err := http.Get(url)
	if err != nil {
		return &StockInfo{}, errors.New("Couldn't fetch stock " + symbol)
	}
	defer resp.Body.Close()

	records, err2 := csv.NewReader(resp.Body).ReadAll()
	if err2 != nil {
		return &StockInfo{}, errors.New("Couldn't read CVS for " + symbol)
	}

	record := records[0]
	floats := make([]float64, len(record)-1)

	for i := 1; i <= len(record)-1; i++ {
		if record[i] == "N/A" {
			floats[i-1] = -1
			continue
		}

		if floats[i-1], err = strconv.ParseFloat(record[i], 64); err != nil {
			errstr := fmt.Sprintf("Couldn't parse float in CVS: %+v", record)
			return &StockInfo{}, errors.New(errstr)
		}
	}

	return &StockInfo{symbol, record[0], floats[0], floats[1], floats[2]}, nil
}

// Causes stock info to be re-fetched and the underlying structure updated
func (me *StockInfo) Update() {
	duplicate, err := NewStockInfo(me.Symbol)
	if err != nil {
		drawDebugText("Couldn't reload existing stock " + me.Symbol)
	}

	me.Name = duplicate.Name
	me.LastTradePrice = duplicate.LastTradePrice
	me.OpeningPrice = duplicate.OpeningPrice
	me.PreviousClosingPrice = duplicate.PreviousClosingPrice
}
