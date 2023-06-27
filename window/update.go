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
	metric := w.MetricScroller.Current()
	metric.Capacity = w.CapScroller.Current()
	w.Chart.Series("serie",
		metric.Queue.Collect()[metric.Queue.Len()-metric.Cap():],
		linechart.SeriesCellOpts(cell.FgColor(w.Color)),
	)
	for _, cell := range w.Cells {
		cell.Update()
	}
}
