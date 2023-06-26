package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/sparkline"
	metric "github.com/vexgratia/termon-go/metric"
	scroller "github.com/vexgratia/termon-go/scroller"
	"github.com/vexgratia/termon-go/update"
)

type Window struct {
	Name    string
	Color   cell.Color
	Metrics []*metric.Metric
	//
	Layout    WindowLayout
	LayoutSet map[WindowLayout]LayoutFunc
	Mode      DisplayMode
	//
	Settings       *button.Button
	Return         *button.Button
	MetricScroller *scroller.Scroller[*metric.Metric]
	Chart          *linechart.LineChart
	Spark          *sparkline.SparkLine
	//
	Updates chan update.Message
}

func New(name string, color cell.Color, metrics []*metric.Metric, updates chan update.Message) *Window {
	window := &Window{
		Name:    name,
		Color:   color,
		Metrics: metrics,

		Layout: WINDOW_DEFAULT,
	}
	window.LayoutSet = map[WindowLayout]LayoutFunc{
		WINDOW_DEFAULT: window.DefaultLayout,
	}
	window.Settings = window.MakeSettingsButton()
	window.Return = window.MakeReturnButton()
	window.Chart = window.MakeChart()
	return window
}
func (w *Window) Opts() []container.Option {
	return w.LayoutSet[w.Layout]()
}
