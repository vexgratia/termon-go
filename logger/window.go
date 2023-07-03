package logger

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
)

type Window struct {
	Color   cell.Color
	Display *text.Text
	Chart   *linechart.LineChart
}

func MakeWindow() *Window {
	window := &Window{}
	window.SetColor(cell.ColorWhite)
	return window
}
func (w *Window) SetColor(color cell.Color) {
	w.Color = color
	w.ResetWidgets()
}
