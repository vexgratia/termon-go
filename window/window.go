package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/scroller"
	"github.com/vexgratia/termon-go/update"
)

var capacities = []uint32{
	500, 200, 100, 10000, 5000, 2000, 1000,
}

type Window struct {
	Name    string
	Color   cell.Color
	Metrics []*metric.Metric
	//
	Layout    WindowLayout
	LayoutSet map[WindowLayout]LayoutFunc
	//
	Settings       *button.Button
	Return         *button.Button
	MetricScroller *scroller.Scroller[*metric.Metric]
	CapScroller    *scroller.Scroller[uint32]
	Chart          *linechart.LineChart
	//
	Updates chan update.Message
}

func New(name string, color cell.Color, metrics []*metric.Metric, updates chan update.Message) *Window {
	window := &Window{
		Name:    name,
		Color:   color,
		Metrics: metrics,

		Layout:  WINDOW_DEFAULT,
		Updates: updates,
	}
	window.LayoutSet = map[WindowLayout]LayoutFunc{
		WINDOW_DEFAULT:  window.DefaultLayout,
		WINDOW_SETTINGS: window.SettingsLayout,
	}
	window.MetricScroller = scroller.New(window.Metrics, window.Color, window.MetricFormat)
	window.CapScroller = scroller.New(capacities, window.Color, window.CapFormat)
	window.Settings = window.MakeSettingsButton()
	window.Return = window.MakeReturnButton()
	window.Chart = window.MakeChart()
	window.Update()
	return window
}
func (w *Window) Opts() []container.Option {
	return w.LayoutSet[w.Layout]()
}
