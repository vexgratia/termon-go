package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
)

func (w *Window) MakeSettingsButton() *button.Button {
	button, _ := button.New(
		"SET",
		func() error {
			w.Layout = w.SettingsLayout
			return nil
		},
		button.Height(2),
		button.FillColor(w.Color),
	)
	return button
}
func (w *Window) MakeReturnButton() *button.Button {
	button, _ := button.New(
		"RETURN",
		func() error {
			w.Layout = w.DefaultLayout
			return nil
		},
		button.Height(2),
		button.FillColor(w.Color),
	)
	return button
}
func (w *Window) MakeChart() *linechart.LineChart {
	chart, _ := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YAxisAdaptive(),
	)
	return chart
}
