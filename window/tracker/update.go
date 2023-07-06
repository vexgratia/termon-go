package tracker

import (
	"time"
)

var tick = 5 * time.Millisecond

// update updates Tracker metrics, templates and widgets.
func (t *Tracker) update() {
	for _, m := range t.data {
		m.Update()
	}
	t.metric.Update()
	t.color.Update()
}

// GetUpdates updates Tracker with current tick.
func (t *Tracker) GetUpdates() {
	for {
		t.update()
		time.Sleep(tick)
	}
}

// relayout updates Tracker container with current opts.
func (t *Tracker) relayout() {
	t.updater.Request(t)
}

// setLayout sets Tracker layout to given LayoutFunc.
//
// Calls relayout.
func (t *Tracker) setLayout(layout LayoutFunc) {
	t.layout = layout
	t.relayout()
}

// reColor colors all Tracker templates, widgets and metrics to current color.
func (t *Tracker) reColor() {
	// // metrics
	// color := t.color.Current()
	// for _, m := range t.data {
	// 	m.SetColor(color)
	// }
	// // widgets
	t.reset()
	// // templates
	// t.color.SetColor(color)
	// t.metric.SetColor(color)
	t.relayout()
}
