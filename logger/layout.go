package logger

import "github.com/mum4k/termdash/container"

func (w *Window) Layout() []container.Option {
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(container.PlaceWidget(w.Display)),
			container.Bottom(container.PlaceWidget(w.Chart)),
			container.SplitPercent(50),
		),
	}
	return opts
}
