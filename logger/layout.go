package logger

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type LayoutFunc func() []container.Option

func (l *Logger) LogLayout() []container.Option {
	layout := []container.Option{}
	layout = append(layout, l.MainOpts()...)
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.Border(linestyle.Round),
				container.BorderTitle(" LOG "),
				container.BorderColor(cell.ColorWhite),
				container.FocusedColor(cell.ColorWhite),
				container.PlaceWidget(l.Display),
			),
			container.Bottom(container.PlaceWidget(l.Spark)),
			container.SplitPercent(85),
		),
	}
	layout = append(layout, opts...)
	return layout
}
func (l *Logger) MainOpts() []container.Option {
	return []container.Option{
		container.ID(l.name),
		container.Border(linestyle.Round),
		container.BorderTitle(" " + l.name + " "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(l.color),
		container.FocusedColor(l.color),
	}
}
