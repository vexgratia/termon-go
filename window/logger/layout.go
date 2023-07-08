package logger

// This file contains the implementation of Logger layouts.

import (
	"strings"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type LayoutFunc func() []container.Option

// initOpts creates opts for Logger container.
func (l *Logger) initOpts() []container.Option {
	return []container.Option{
		container.ID(l.name),
		container.Border(linestyle.Round),
		container.BorderTitle(" " + strings.ToUpper(l.name) + " "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(l.Color()),
		container.FocusedColor(l.Color()),
	}
}

// logLayout contains settings button, display and spark.
func (l *Logger) logLayout() []container.Option {
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.Border(linestyle.Round),
				container.BorderTitle(" LOG "),
				container.BorderColor(cell.ColorWhite),
				container.FocusedColor(cell.ColorWhite),
				container.PlaceWidget(l.display),
			),
			container.Bottom(container.PlaceWidget(l.spark)),
			container.SplitPercent(85),
		),
	}
	return opts
}
