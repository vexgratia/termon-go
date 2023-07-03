package logger

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
)

func (w *Window) ResetWidgets() {
	w.Chart = w.MakeChart()
	w.Display = w.MakeDisplay()
}
func (w *Window) MakeChart() *linechart.LineChart {
	chart, _ := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorWhite)),
	)
	return chart
}
func (w *Window) MakeDisplay() *text.Text {
	text, _ := text.New(
		text.WrapAtWords(),
		text.DisableScrolling(),
	)
	return text
}
