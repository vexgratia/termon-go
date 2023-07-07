package selecter

// This file contains the implementation of Selecter layouts.

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/vexgratia/termon-go/window"
)

type LayoutFunc func() []container.Option

// Opts returns Selecter layout.
func (s *Selecter) Opts() []container.Option {
	layout := []container.Option{
		container.ID(s.Name()),
		container.Border(linestyle.Round),
		container.BorderTitle(fmt.Sprintf(" WINDOW %d ", s.id)),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
	}
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, s.makeContainer(0)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(1)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(2)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(3)),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, s.makeContainer(4)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(5)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(6)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(7)),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, s.makeContainer(8)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(9)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(10)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(11)),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, s.makeContainer(12)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(13)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(14)),
			grid.ColWidthPercWithOpts(25, s.makeContainer(15)),
		),
	)
	opts, _ := builder.Build()
	layout = append(layout, opts...)
	return layout
}

// makeContainer creates container for Window button using index.
//
// If index is not represented in Windows array, returns empty container.
func (s *Selecter) makeContainer(index int) []container.Option {
	if index >= len(s.windows) || index < 0 {
		return []container.Option{}
	}
	return []container.Option{container.PlaceWidget(s.makeButton(s.windows[index]))}
}

// makeButton creates button that swaps Selecter with given Window.
func (s *Selecter) makeButton(w window.Window) *button.Button {
	b, _ := button.New(
		w.Name(),
		s.windowSetFunc(w),
		button.Height(2),
		button.Width(8),
		button.FillColor(w.Color()),
	)
	return b
}
