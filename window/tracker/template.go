package tracker

// This file contains the implementation of Tracker templates.

import (
	"github.com/mum4k/termdash/cell"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/palette"
	"github.com/vexgratia/termon-go/template/scroller"
)

// makeMetricScroller creates a Scroller for Tracker Metrics.
func (t *Tracker) makeMetricScroller() *scroller.Scroller[*metric.Metric] {
	metric := scroller.New[*metric.Metric]()
	metric.SetScrollFunc(t.relayout)
	metric.SetFormatter(metricFormat)
	return metric
}

// makeColorScroller creates a Scroller for Tracker colors.
func (t *Tracker) makeColorScroller() *scroller.Scroller[cell.Color] {
	color := scroller.New[cell.Color]()
	color.SetScrollFunc(t.reColor)
	color.SetFormatter(colorFormat)
	color.Add(palette.All...)
	return color
}
