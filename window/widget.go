package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/vexgratia/termon-go/metric"
)

func (w *Window) MakeSettingsButton() *button.Button {
	button, _ := button.New(
		"SET",
		func() error {
			w.Layout = w.SettingsLayout
			w.UpdateLayout()
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
			w.UpdateLayout()
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
			w.UpdateLayout()
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(w.Color),
	)
	return button
}
func (w *Window) MakeChart(metric *metric.Metric) *linechart.LineChart {
	chart, _ := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YAxisFormattedValues(linechart.ValueFormatter(
			metric.Format,
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
