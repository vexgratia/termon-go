package tracker

import (
	"time"
)

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
		time.Sleep(t.Tick())
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
//
// Calls relayout.
func (t *Tracker) reColor() {
	t.reset()
	t.relayout()
}
