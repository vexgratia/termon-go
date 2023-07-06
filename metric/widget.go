package metric

// This file contains the implementation of Metric widgets.

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
)

// reset recreates all Metric widgets.
//
// Blocks Metric mutex to avoid data race.
func (m *Metric) reset() {
	m.mu.Lock()
	m.display = m.makeDisplay()
	m.chart = m.makeChart()
	m.mu.Unlock()
}

// makeDisplay creates Metric display.
func (m *Metric) makeDisplay() *text.Text {
	text, _ := text.New(
		text.WrapAtRunes(),
		text.DisableScrolling(),
	)
	return text
}

// makeChart creates Metric chart.
func (m *Metric) makeChart() *linechart.LineChart {
	chart, _ := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YAxisFormattedValues(linechart.ValueFormatter(
			m.formatter,
		)),
		linechart.YAxisAdaptive(),
	)
	return chart
}
