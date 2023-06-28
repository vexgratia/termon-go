package window

import (
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/vexgratia/termon-go/cache"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/scroller"
	"github.com/vexgratia/termon-go/update"
)

var capacities = []uint32{
	2000, 5000, 10000, 10, 20, 50, 100, 200, 500, 1000,
}
var ticks = []time.Duration{
	200, 500, 1000, 2000, 5, 10, 20, 50, 100,
}

type Window struct {
	Name    string
	Color   cell.Color
	Metrics []*metric.Metric
	//
	Layout  LayoutFunc
	ShowAvg *scroller.Scroller[bool]
	ShowMax *scroller.Scroller[bool]
	//
	Settings       *button.Button
	MetricScroller *scroller.Scroller[*metric.Metric]
	CapScroller    *scroller.Scroller[uint32]
	TickScroller   *scroller.Scroller[time.Duration]
	//
	Chart       *linechart.LineChart
	Cells       []*Cell
	ChartButton *button.Button
	CellButton  *button.Button
	//
	Cache   *cache.Cache
	Updates chan update.Message
}

func New(name string, color cell.Color, metrics []string, updates chan update.Message) *Window {
	window := &Window{
		Name:    name,
		Color:   color,
		Updates: updates,
	}
	//
	window.ShowAvg = scroller.New([]bool{false, true}, window.Color, window.AvgScrollerFormat, func() {})
	window.ShowMax = scroller.New([]bool{false, true}, window.Color, window.MaxScrollerFormat, func() {})
	//
	window.Settings = window.MakeSettingsButton()
	window.ChartButton = window.MakeChartButton()
	window.CellButton = window.MakeCellButton()
	//
	window.TickScroller = scroller.New(ticks, window.Color, window.TickScrollerFormat, func() {})
	window.Cache = cache.New(metrics)
	window.Metrics = window.Cache.GetMetrics()
	//
	window.MetricScroller = scroller.New(window.Metrics, window.Color, window.MetricScrollerFormat,
		func() {
			window.Chart = window.MakeChart(window.MetricScroller.Current())
			window.UpdateLayout()
		})
	window.CapScroller = scroller.New(capacities, window.Color, window.CapScrollerFormat, func() {})
	window.Chart = window.MakeChart(window.MetricScroller.Current())
	window.Cells = window.MakeCells()
	//
	window.Layout = window.ChartLayout
	go window.GetUpdates()
	return window
}
func (w *Window) Opts() []container.Option {
	return w.Layout()
}
