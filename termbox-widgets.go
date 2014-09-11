package main

import (
	"fmt"
	"math"
	"time"

	"github.com/nsf/termbox-go"
)

type Widget interface {
	Size() (int, int)
	Draw(x, y int)
	Update()
}

/* Sleeps a bit, then calls termbox.Flush */
func sleepFlush() {
	time.Sleep(500 * time.Millisecond)
	termbox.Flush()
}

/* Draws the specified box to the termbox buffer */
func drawBox(x, y, width, height int) {
	fg, bg := termbox.ColorWhite, termbox.ColorDefault
	width, height = width-1, height-1

	// Corners
	termbox.SetCell(x, y, '┌', fg, bg)
	termbox.SetCell(x+width, y, '┐', fg, bg)
	termbox.SetCell(x+width, y+height, '┘', fg, bg)
	termbox.SetCell(x, y+height, '└', fg, bg)

	// Top & Bottom Lines
	for i := 1; i < width; i++ {
		termbox.SetCell(x+i, y, '─', fg, bg)
		termbox.SetCell(x+i, y+height, '─', fg, bg)
	}

	// Right & Left Lines
	for i := 1; i < height; i++ {
		termbox.SetCell(x, y+i, '│', fg, bg)
		termbox.SetCell(x+width, y+i, '│', fg, bg)
	}

	termbox.Flush()
}

func drawText(x, y int, text string) {
	fg, bg := termbox.ColorWhite, termbox.ColorDefault

	for index, char := range text {
		termbox.SetCell(x+index, y, char, fg, bg)
	}

	termbox.Flush()
}

func drawDebugText(text string) {
	pageWidth, pageHeight := termbox.Size()
	fg, bg := termbox.ColorWhite, termbox.ColorDefault

	for i := 0; i < pageWidth; i++ {
		termbox.SetCell(i, pageHeight-1, ' ', fg, bg)
	}

	drawText(1, pageHeight-1, text)
}

type StockInfoWidget struct {
	StockInfo *StockInfo
}

func NewStockInfoWidget(stockInfo *StockInfo) *StockInfoWidget {
	return &StockInfoWidget{stockInfo}
}

func (me *StockInfoWidget) Size() (int, int) {
	width := math.Max(10, float64(len(me.StockInfo.Name)+3))
	height := 4

	return int(width), height
}

func (me *StockInfoWidget) Draw(x, y int) {
	width, height := me.Size()
	openPriceStr := fmt.Sprintf("%v", me.StockInfo.LastTradePrice)

	drawBox(x, y, width, height)
	drawText(x+2, y+1, me.StockInfo.Name)
	drawText(x+2, y+2, openPriceStr)
}

func (me *StockInfoWidget) Update() {
	me.StockInfo.Update()
}
