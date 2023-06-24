package window

import "github.com/mum4k/termdash/widgetapi"

type WindowMode int

const (
	LINECHART WindowMode = iota
	SPARKLINE
)

var WindowModes = []WindowMode{
	LINECHART,
	SPARKLINE,
}

func (w *Window) ModeWidget() widgetapi.Widget {
	switch w.ModeScroller.Current() {
	case LINECHART:
		return w.Chart
	case SPARKLINE:
		return w.Spark
	}
	return nil
}
