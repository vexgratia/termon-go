package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/vexgratia/termon-go"
	"github.com/vexgratia/termon-go/scroller"
)

type LayoutFunc func() []container.Option

type Window struct {
	Name    string
	Color   cell.Color
	Metrics []*termon.Metric
	Layout  LayoutFunc
	//
	Settings       *button.Button
	Return         *button.Button
	MetricScroller *scroller.Scroller[*termon.Metric]
	Chart          *linechart.LineChart
	Spark          *sparkline.SparkLine
	//
	ModeScroller *scroller.Scroller[WindowMode]
}

func MakeWindow(name string, color cell.Color, metrics []*termon.Metric) *Window {
	window := &Window{
		Name:    name,
		Color:   color,
		Metrics: metrics,
	}
	window.Settings = window.MakeSettingsButton()
	window.Return = window.MakeReturnButton()
	window.MetricScroller = scroller.MakeScroller(window.Metrics, window.Color, scroller.MetricDisplayFormatter)
	window.ModeScroller = scroller.MakeScroller(WindowModes, window.Color, ModeFormatter)
	window.Chart = window.MakeChart()
	return window
}
