package selecter

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
)

type LayoutFunc func() []container.Option

func (s *Selecter) Opts() []container.Option {
	layout := []container.Option{
		container.ID(s.Name()),
		container.Border(linestyle.Round),
		container.BorderTitle(fmt.Sprintf(" WINDOW %d ", s.ID)),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
	}
	containers := s.MakeContainers()
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, containers[0]),
			grid.ColWidthPercWithOpts(25, containers[1]),
			grid.ColWidthPercWithOpts(25, containers[2]),
			grid.ColWidthPercWithOpts(25, containers[3]),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, containers[4]),
			grid.ColWidthPercWithOpts(25, containers[5]),
			grid.ColWidthPercWithOpts(25, containers[6]),
			grid.ColWidthPercWithOpts(25, containers[7]),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, containers[8]),
			grid.ColWidthPercWithOpts(25, containers[9]),
			grid.ColWidthPercWithOpts(25, containers[10]),
			grid.ColWidthPercWithOpts(25, containers[11]),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, containers[12]),
			grid.ColWidthPercWithOpts(25, containers[13]),
			grid.ColWidthPercWithOpts(25, containers[14]),
			grid.ColWidthPercWithOpts(25, containers[15]),
		),
	)
	opts, _ := builder.Build()
	layout = append(layout, opts...)
	return layout
}
