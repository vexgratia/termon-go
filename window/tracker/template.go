package tracker

// This file contains the implementation of Tracker templates.

import (
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/palette"
	"github.com/vexgratia/termon-go/template/scroller"
)

var ticks = []time.Duration{
	5 * time.Millisecond,
	10 * time.Millisecond,
	100 * time.Millisecond,
	500 * time.Millisecond,
	1000 * time.Millisecond,
}

// makeMetricScroller creates a Scroller for Tracker Metrics.
func (t *Tracker) makeMetricScroller() *scroller.Scroller[*metric.Metric] {
	metric := scroller.New[*metric.Metric]()
	metric.SetScrollFunc(t.relayout)
	metric.SetFormatter(t.metricFormat)
	return metric
}

// makeColorScroller creates a Scroller for Tracker colors.
func (t *Tracker) makeColorScroller() *scroller.Scroller[cell.Color] {
	color := scroller.New[cell.Color]()
	color.SetScrollFunc(t.reColor)
	color.SetFormatter(t.colorFormat)
	color.Add(palette.All...)
	return color
}

// makeTickScroller creates a Scroller for Tracker tick.
func (t *Tracker) makeTickScroller() *scroller.Scroller[time.Duration] {
	tick := scroller.New[time.Duration]()
	tick.SetFormatter(t.tickFormat)
	tick.Add(ticks...)
	return tick
}
