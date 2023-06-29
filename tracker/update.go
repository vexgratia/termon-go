package tracker

import (
	"time"

	"github.com/vexgratia/termon-go/update"
)

func (t *Tracker) Connect(signal chan update.Message) {
	t.Signal = signal
}
func (t *Tracker) Relayout() {
	t.Signal <- update.Message{
		Name: t.name,
		Opts: t.Opts(),
	}
}

func (t *Tracker) SetLayout(layout LayoutFunc) {
	t.Layout = layout
	t.Relayout()
}
func (t *Tracker) GetUpdates() {
	for {
		t.MetricScroller.Update()
		for _, metric := range t.Metrics {
			metric.Update()
		}
		time.Sleep(time.Millisecond * 5)
	}
}
