package tracker

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/scroller"
	"github.com/vexgratia/termon-go/update"
)

type Tracker struct {
	name           string
	Color          cell.Color
	Layout         LayoutFunc
	Metrics        []*metric.Metric
	MetricScroller *scroller.Scroller[*metric.Metric]
	Settings       *button.Button
	Signal         chan update.Message
}

func New(name string) *Tracker {
	tracker := &Tracker{
		name:           name,
		MetricScroller: scroller.New[*metric.Metric](),
	}
	tracker.MetricScroller.SetScrollFunc(tracker.Relayout)
	tracker.SetColor(cell.ColorWhite)
	return tracker
}
func (t *Tracker) Name() string {
	return t.name
}
func (t *Tracker) Add(metrics ...*metric.Metric) {
	for _, metric := range metrics {
		t.MetricScroller.Add(metric)
		t.Metrics = append(t.Metrics, metric)
	}
}

func (t *Tracker) SetColor(color cell.Color) {
	t.Color = color
	t.MetricScroller.SetColor(color)
	t.ResetWidgets()
}
