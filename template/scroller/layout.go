package scroller

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
)

// Layout returns Scroller opts.
func (s *Scroller[T]) Layout() []container.Option {
	builder := grid.New()
	builder.Add(
		grid.ColWidthPerc(20,
			grid.Widget(s.prev,
				container.Border(linestyle.None),
			),
		),
	)
	builder.Add(
		grid.ColWidthPerc(60,
			grid.Widget(s.display,
				container.Border(linestyle.Light)),
		),
	)
	builder.Add(
		grid.ColWidthPerc(20,
			grid.Widget(s.next,
				container.Border(linestyle.None),
			),
		),
	)
	opts, _ := builder.Build()
	return opts
}
