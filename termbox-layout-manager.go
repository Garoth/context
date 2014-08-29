package main

import (
	"math"

	"github.com/nsf/termbox-go"
)

type InlineBlockLayout struct {
	Widgets []Widget
}

func NewInlineBlockLayout() *InlineBlockLayout {
	return &InlineBlockLayout{make([]Widget, 0)}
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
