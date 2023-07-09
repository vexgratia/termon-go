package logger

// This file contains the implementation of Logger layouts.

import (
	"strings"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
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

// settingsLayout contains tickScroller, colorScroller, cell and chart buttons.
func (l *Logger) settingsLayout() []container.Option {
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(25,
			grid.Widget(l.log),
		),
	)
	builder.Add(
		grid.RowHeightPercWithOpts(25, l.color.Layout()),
	)
	opts, _ := builder.Build()
	return opts
}

// logLayout contains settings button, display and spark.
func (l *Logger) logLayout() []container.Option {
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(
						container.PlaceWidget(l.settings),
					),
					container.Right(container.PlaceWidget(l.spark)),
					container.SplitPercent(25),
				),
			),
			container.Bottom(
				container.PlaceWidget(l.display),
			),
			container.SplitPercent(30),
		),
	}
	return opts
}
