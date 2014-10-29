package main

import "github.com/nsf/termbox-go"

func refreshWidgets(desiredStocks []string, layoutManager *InlineBlockLayout) {
	layoutManager.Clear()
	layoutManager.Redraw()

	for i := 0; i < len(desiredStocks); i++ {
		stock := desiredStocks[i]
		stockInfo, err := NewStockInfo(stock)
		if err != nil {
			panic(err)
		}

		widget := NewStockInfoWidget(stockInfo)
		layoutManager.Add(widget)

		if i != len(desiredStocks)-1 {
			divider := NewDividerWidget(4)
			layoutManager.Add(divider)
		}
	}
}

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
			// KEY EVENTS
			if ev.Key == termbox.KeyEsc {
				break
			} else if ev.Key == termbox.KeySpace {
				refreshWidgets(desiredStocks, layoutManager)
			}

		} else if ev.Type == termbox.EventResize {
			// RESIZE EVENT
			refreshWidgets(desiredStocks, layoutManager)

		} else if ev.Type == termbox.EventMouse {
			// MOUSE EVENT

		} else if ev.Type == termbox.EventError {
			// ERROR EVENT
			panic(ev.Err)
		}
	}
}
