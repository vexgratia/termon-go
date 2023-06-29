package metric

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
)

func (m *Metric) ResetWidgets() {
	m.Chart = m.MakeChart()
	m.Display = m.MakeDisplay()
}
func (m *Metric) MakeDisplay() *text.Text {
	text, _ := text.New(
		text.WrapAtRunes(),
		text.DisableScrolling(),
	)
	return text
}
func (m *Metric) MakeChart() *linechart.LineChart {
	chart, _ := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YAxisFormattedValues(linechart.ValueFormatter(
			m.Format,
		)),
		linechart.YAxisAdaptive(),
	)
	return chart
}
