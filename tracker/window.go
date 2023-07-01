package tracker

import (
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
)

func (t *Tracker) Name() string {
	return t.name
}
func (t *Tracker) Color() cell.Color {
	return t.color
}
func (t *Tracker) Lock() {
	t.mu.Lock()
}
func (t *Tracker) Unlock() {
	t.mu.Unlock()
}
func (t *Tracker) Opts() []container.Option {
	return t.Layout()
}
func (t *Tracker) Run() {
	for {
		t.MetricScroller.Update()
		for _, metric := range t.Metrics {
			metric.Update()
		}
		time.Sleep(time.Millisecond * 5)
	}
}
