package logger

import "time"

var tick = 5 * time.Millisecond

// update updates Logger templates and widgets.
func (l *Logger) update() {
	l.updateWidgets()
	l.color.Update()
}

// updateWidgets updates Logger widgets.
func (l *Logger) updateWidgets() {
	// spark
	l.spark.Add([]int{l.data.Len()})
	l.data.Clear()
}

// GetUpdates updates Logger with current tick.
func (l *Logger) GetUpdates() {
	for {
		l.update()
		time.Sleep(tick)
	}
}

// relayout updates Logger container with current opts.
func (l *Logger) relayout() {
	l.updater.Request(l)
}

// setLayout sets Logger layout to given LayoutFunc.
//
// Calls relayout.
func (l *Logger) setLayout(layout LayoutFunc) {
	l.layout = layout
	l.relayout()
}

// reColor colors all Logger templates and widgets to current color.
//
// Calls relayout.
func (l *Logger) reColor() {
	l.reset()
	l.relayout()
}
