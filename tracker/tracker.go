package tracker

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/palette"
	"github.com/vexgratia/termon-go/scroller"
	"github.com/vexgratia/termon-go/updater"
)

type Tracker struct {
	name           string
	color          cell.Color
	Layout         LayoutFunc
	Metrics        []*metric.Metric
	MetricScroller *scroller.Scroller[*metric.Metric]
	ColorScroller  *scroller.Scroller[cell.Color]
	Settings       *button.Button
	Chart          *button.Button
	Cell           *button.Button
	Updater        *updater.Updater
}

func New(name string, updater *updater.Updater) *Tracker {
	tracker := &Tracker{
		name:           name,
		MetricScroller: scroller.New[*metric.Metric](),
		ColorScroller:  scroller.New[cell.Color](),
		Updater:        updater,
	}
	tracker.Layout = tracker.ChartLayout
	//
	tracker.ColorScroller.Add(palette.All...)
	tracker.SetColor(tracker.ColorScroller.Current())
	tracker.ColorScroller.SetScrollFunc(
		func() {
			tracker.SetColor(tracker.ColorScroller.Current())
			tracker.ColorScroller.Update()
			tracker.Relayout()
		},
	)
	tracker.ColorScroller.SetFormatter(ColorFormatter)
	//
	tracker.MetricScroller.SetScrollFunc(tracker.Relayout)
	tracker.MetricScroller.SetFormatter(MetricFormatter)
	//
	return tracker
}

func (t *Tracker) Add(metrics ...*metric.Metric) {
	for _, metric := range metrics {
		t.MetricScroller.Add(metric)
		t.Metrics = append(t.Metrics, metric)
	}
}

func (t *Tracker) SetColor(color cell.Color) {
	t.color = color
	t.MetricScroller.SetColor(color)
	t.ColorScroller.SetColor(color)
	for _, m := range t.Metrics {
		m.SetColor(color)
	}
	t.ResetWidgets()
}
