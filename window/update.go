package window

import (
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/vexgratia/termon-go/update"
)

func (w *Window) UpdateLayout() {
	w.Updates <- update.Message{
		Name: w.Name,
		Opts: w.Opts(),
	}
}
func (w *Window) Update() {
	w.MetricScroller.Update()
	w.CapScroller.Update()
	w.TickScroller.Update()
	metric := w.MetricScroller.Current()
	metric.Capacity = w.CapScroller.Current()
	w.Chart.Series("data",
		metric.Data.Collect(),
		linechart.SeriesCellOpts(cell.FgColor(w.Color)),
	)
	for _, c := range w.Cells {
		c.Update()
	}
}
func (w *Window) GetUpdates() {
	for {
		w.Cache.Update()
		w.Update()
		time.Sleep(w.TickScroller.Current() * 1000000)
	}
}
