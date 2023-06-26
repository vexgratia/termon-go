package window

import "github.com/mum4k/termdash/widgetapi"

type DisplayMode int

const (
	LINECHART DisplayMode = iota
	SPARKLINE
)

func (w *Window) DisplayWidget() widgetapi.Widget {
	switch w.Mode {
	case LINECHART:
		return w.Chart
	case SPARKLINE:
		return w.Spark
	}
	return nil
}
