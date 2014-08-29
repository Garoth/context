package main

import "github.com/nsf/termbox-go"

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	layoutManager := NewInlineBlockLayout()
	desiredStocks := []string{"BTCUSD=X", "TSLA", "GOOG", "AAPL"}

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc {
				break
			} else if ev.Key == termbox.KeySpace {
				for _, stock := range desiredStocks {
					stockInfo, err := GetStockInfo(stock)
					if err != nil {
						panic(err)
					}

					widget := NewStockInfoWidget(stockInfo)
					layoutManager.Add(widget)
				}
			}
		} else if ev.Type == termbox.EventError {
			panic(ev.Err)
		}
	}
}
