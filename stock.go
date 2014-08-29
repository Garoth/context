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
	Name                 string
	LastTradePrice       float64
	OpeningPrice         float64
	PreviousClosingPrice float64
}

/* Creates a new StockInfo from an array of values that are ordered the same
 * as the fields of the StockInfo struct. */
func NewStockInfo(record []string) (*StockInfo, error) {
	floats := make([]float64, len(record)-1)

	var err error
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

	return &StockInfo{record[0], floats[0], floats[1], floats[2]}, nil
}

func GetStockInfo(symbol string) (*StockInfo, error) {
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

	return NewStockInfo(records[0])
}
