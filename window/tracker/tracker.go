package tracker

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/palette"
	"github.com/vexgratia/termon-go/template/scroller"
	"github.com/vexgratia/termon-go/updater"
)

type Tracker struct {
	name     string
	Layout   LayoutFunc
	Metrics  []*metric.Metric
	metric   *scroller.Scroller[*metric.Metric]
	color    *scroller.Scroller[cell.Color]
	Settings *button.Button
	Chart    *button.Button
	Cell     *button.Button
	Updater  *updater.Updater
}

func New(name string, updater *updater.Updater) *Tracker {
	tracker := &Tracker{
		name:    name,
		metric:  scroller.New[*metric.Metric](),
		color:   scroller.New[cell.Color](),
		Updater: updater,
	}
	tracker.Layout = tracker.ChartLayout
	//
	tracker.color.Add(palette.All...)
	tracker.color.SetScrollFunc(
		func() {
			tracker.SetColor(tracker.color.Current())
			tracker.color.Update()
			tracker.Relayout()
		},
	)
	tracker.color.SetFormatter(ColorFormatter)
	//
	tracker.metric.SetScrollFunc(tracker.Relayout)
	tracker.metric.SetFormatter(MetricFormatter)
	//
	return tracker
}

func (t *Tracker) Add(metrics ...*metric.Metric) {
	for _, metric := range metrics {
		t.metric.Add(metric)
		t.Metrics = append(t.Metrics, metric)
	}
}

func (t *Tracker) SetColor(color cell.Color) {
	t.metric.SetColor(color)
	t.color.SetColor(color)
	for _, m := range t.Metrics {
		m.SetColor(color)
	}
	t.ResetWidgets()
}
