package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
)

func (w *Window) Update() {
	w.MetricScroller.Update()
	w.Chart.Series("serie", w.MetricScroller.Current().Queue.Collect(), linechart.SeriesCellOpts(cell.FgColor(w.Color)))
}
