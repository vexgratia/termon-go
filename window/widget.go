package window

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/sparkline"
)

func (w *Window) MakeSettingsButton() *button.Button {
	button, _ := button.New(
		"SET",
		func() error {
			w.Layout = w.SettingsLayout
			w.Update()
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(w.Color),
	)
	return button
}

func (w *Window) MakeChartButton() *button.Button {
	button, _ := button.New(
		"CHART",
		func() error {
			w.Layout = w.ChartLayout
			w.Update()
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(w.Color),
	)
	return button
}

func (w *Window) MakeCellButton() *button.Button {
	button, _ := button.New(
		"CELL",
		func() error {
			w.Layout = w.CellLayout
			w.Update()
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(w.Color),
	)
	return button
}
func (w *Window) MakeChart() *linechart.LineChart {
	chart, _ := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YAxisFormattedValues(linechart.ValueFormatter(
			func(value float64) string {
				return fmt.Sprintf("%5.1G", value)
			},
		)),
		linechart.YAxisAdaptive(),
	)
	return chart
}
func (w *Window) MakeSpark() *sparkline.SparkLine {
	spark, _ := sparkline.New(
		sparkline.Color(w.Color),
	)
	return spark
}
