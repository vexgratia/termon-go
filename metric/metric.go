package metric

import (
	"runtime/metrics"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
	queue "github.com/vexgratia/collection-go/generic/queue"
)

type Metric struct {
	name   string
	Color  cell.Color
	sample []metrics.Sample
	//
	Tag  string
	Type MetricType
	//
	Data *queue.Queue[float64]
	//
	Display *text.Text
	Chart   *linechart.LineChart
	//
	Format func(value float64) string
}

func New(name string) *Metric {
	metric := &Metric{
		name:   name,
		sample: []metrics.Sample{{Name: name}},
		Data:   queue.New[float64](),
	}
	metric.Tag, metric.Type = ParseTable[name].Tag, ParseTable[name].Type
	metric.UpdateData()
	return metric
}
func (m *Metric) Name() string {
	return m.name
}
func (m *Metric) Current() float64 {
	return m.Data.Peek()
}
func (m *Metric) SetColor(color cell.Color) {
	m.Color = color
	m.ResetWidgets()
}
