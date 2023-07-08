package metric

import (
	"log"
	"runtime/metrics"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
)

// SetColor sets Metric color to given color.
func (m *Metric) SetColor(color cell.Color) {
	m.mu.Lock()
	m.color = color
	m.mu.Unlock()
}

// update updates Metric data and widgets.
func (m *Metric) Update() {
	m.updateData()
	m.updateWidgets()
}

// updateWidgets updates Metric widgets.
func (m *Metric) updateWidgets() {
	// display
	m.display.Reset()
	for _, text := range m.displayFormat() {
		m.display.Write(text.Text, text.Opts...)
	}
	// chart
	m.chart.Series(
		"data", m.data.Collect(),
		linechart.SeriesCellOpts(cell.FgColor(m.color)))
}

// updateData enqueues new data from Golang runtime.
//
// Dequeues if queue exceeds capacity.
func (m *Metric) updateData() {
	var data float64
	sample := []metrics.Sample{{Name: m.name}}
	metrics.Read(sample)
	value := sample[0].Value
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
	m.current = data
	m.data.Enqueue(data)
	if m.data.Len() >= maxCap {
		m.data.Dequeue()
	}
}

// medianBucket is a helper function from Golang runtime/metrics example.
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
