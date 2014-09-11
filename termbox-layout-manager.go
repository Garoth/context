package main

import (
	"math"
	"time"

	"github.com/nsf/termbox-go"
)

type InlineBlockLayout struct {
	Widgets []Widget
}

func NewInlineBlockLayout() *InlineBlockLayout {
	me := &InlineBlockLayout{make([]Widget, 0)}

	go func() {
		for {
			time.Sleep(time.Second * 30)
			for _, widget := range me.Widgets {
				time.Sleep(time.Millisecond * 100)
				widget.Update()
			}
			me.Redraw()
		}
	}()

	return me
}

func (me *InlineBlockLayout) Add(widget Widget) {
	// FIXME there's some bug with widget heights and new line alignment
	pageWidth, _ := termbox.Size()
	x, y, currentLineHeight := 0, 0, 0

	for _, widgetElement := range me.Widgets {
		widgetWidth, widgetHeight := widgetElement.Size()

		x += widgetWidth

		// Keeping track of tallest element in current line
		currentLineHeight = int(math.Max(float64(currentLineHeight),
			float64(widgetHeight)))

		// We've hit a line wrap
		if x >= pageWidth {
			x = widgetWidth
			y += currentLineHeight
			currentLineHeight = 0
		}
	}

	widgetWidth, _ := widget.Size()

	if x+widgetWidth >= pageWidth {
		x = 0
		y += currentLineHeight
	}

	widget.Draw(x, y)

	me.Widgets = append(me.Widgets, widget)
}

func (me *InlineBlockLayout) Redraw() {
	pageWidth, _ := termbox.Size()
	x, y, currentLineHeight := 0, 0, 0

	for _, widgetElement := range me.Widgets {
		widgetWidth, widgetHeight := widgetElement.Size()

		widgetElement.Draw(x, y)
		x += widgetWidth

		// Keeping track of tallest element in current line
		currentLineHeight = int(math.Max(float64(currentLineHeight),
			float64(widgetHeight)))

		// We've hit a line wrap
		if x >= pageWidth {
			x = widgetWidth
			y += currentLineHeight
			currentLineHeight = 0
		}
	}
}
