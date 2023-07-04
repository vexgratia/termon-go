package metric

import (
	"log"
	"runtime/metrics"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
)

var maxCap = 1000

func (m *Metric) Update() {
	m.UpdateData()
	m.UpdateWidgets()
}
func (m *Metric) UpdateWidgets() {
	m.UpdateChart()
	m.UpdateDisplay()
}
func (m *Metric) UpdateDisplay() {
	m.Display.Reset()
	textChunk := m.DisplayFormat()
	for _, unit := range textChunk {
		m.Display.Write(unit.Text, unit.Opts...)
	}
}
func (m *Metric) UpdateChart() {
	m.Chart.Series(
		"data", m.Data.Collect(),
		linechart.SeriesCellOpts(cell.FgColor(m.Color)))
}
func (m *Metric) UpdateData() {
	var data float64
	metrics.Read(m.sample)
	value := m.sample[0].Value
	switch value.Kind() {
	case metrics.KindUint64:
		data = float64(value.Uint64())
	case metrics.KindFloat64:
		data = value.Float64()
	case metrics.KindFloat64Histogram:
		data = medianBucket(value.Float64Histogram())
	case metrics.KindBad:
		log.Fatalf("BAD METRIC KIND")
	default:

	}
	m.Current = data
	m.Data.Enqueue(data)
	if m.Data.Len() >= maxCap {
		m.Data.Dequeue()
	}
}
func medianBucket(h *metrics.Float64Histogram) float64 {
	total := uint64(0)
	for _, count := range h.Counts {
		total += count
	}
	thresh := total / 2
	total = 0
	for i, count := range h.Counts {
		total += count
		if total >= thresh {
			return h.Buckets[i]
		}
	}
	panic("should not happen")
}
