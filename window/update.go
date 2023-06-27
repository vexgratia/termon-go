package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/vexgratia/termon-go/update"
)

func (w *Window) Update() {
	w.Updates <- update.Message{
		Name: w.Name,
		Opts: w.Opts(),
	}
}
func (w *Window) UpdateWidgets() {
	w.MetricScroller.Update()
	w.CapScroller.Update()
	w.MetricScroller.Current().Capacity = w.CapScroller.Current()
	w.Chart.Series("serie",
		w.MetricScroller.Current().Queue.Collect()[w.MetricScroller.Current().Queue.Len()-w.MetricScroller.Current().Cap():],
		linechart.SeriesCellOpts(cell.FgColor(w.Color)))
}
