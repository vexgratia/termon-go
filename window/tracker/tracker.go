package tracker

// This file contains the implementation of Tracker and its basic methods.

import (
	"sync"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/template/scroller"
	"github.com/vexgratia/termon-go/updater"
)

// A Tracker is a Window type that displays Metric data.
type Tracker struct {
	// general
	name   string
	data   []*metric.Metric
	layout LayoutFunc
	// sync
	mu *sync.Mutex
	// templates
	metric *scroller.Scroller[*metric.Metric]
	tick   *scroller.Scroller[time.Duration]
	color  *scroller.Scroller[cell.Color]
	// widgets
	settings *button.Button
	chart    *button.Button
	cell     *button.Button
	// external
	updater *updater.Updater
}

// New creates a Tracker based on Updater.
func New(name string, updater *updater.Updater) *Tracker {
	// general, external and sync
	t := &Tracker{
		name:    name,
		mu:      &sync.Mutex{},
		updater: updater,
	}
	t.layout = t.settingsLayout
	// templates
	t.metric = t.makeMetricScroller()
	t.color = t.makeColorScroller()
	t.tick = t.makeTickScroller()
	// widgets
	t.reset()
	return t
}

// Add appends Metric to Tracker.
func (t *Tracker) Add(metrics ...*metric.Metric) {
	for _, metric := range metrics {
		t.metric.Add(metric)
		t.data = append(t.data, metric)
	}
}

// Name returns Tracker name.
func (t *Tracker) Name() string {
	return t.name
}

// Color returns current Tracker color.
func (t *Tracker) Color() cell.Color {
	return t.color.Current()
}

// Color returns current Tracker color.
func (t *Tracker) Tick() time.Duration {
	return t.tick.Current()
}

// Opts configurates and returns opts based on current LayoutFunc.
func (t *Tracker) Opts() []container.Option {
	opts := t.initOpts()
	opts = append(opts, t.layout()...)
	return opts
}
