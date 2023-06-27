package window

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
)

type WindowLayout int

type LayoutFunc func() []container.Option

const (
	WINDOW_DEFAULT WindowLayout = iota
	WINDOW_SETTINGS
	WINDOW_CELL
)

func (w *Window) DefaultLayout() []container.Option {
	return []container.Option{
		container.ID(w.Name),
		container.Border(linestyle.Round),
		container.BorderTitle(w.Name),
		container.BorderTitleAlignCenter(),

		container.BorderColor(w.Color),
		container.FocusedColor(w.Color),
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(
						container.PlaceWidget(w.Settings),
					),
					container.Right(w.MetricScroller.Layout()...),
					container.SplitPercent(20),
				),
			),
			container.Bottom(
				container.PlaceWidget(w.Chart),
			),
			container.SplitPercent(30),
		),
	}
}
func (w *Window) SettingsLayout() []container.Option {
	return []container.Option{
		container.ID(w.Name),
		container.Border(linestyle.Round),
		container.BorderTitle("SETTINGS"),
		container.SplitHorizontal(
			container.Top(
				container.PlaceWidget(w.Return),
			),
			container.Bottom(
				w.CapScroller.Layout()...,
			),
		),
	}
}
func (w *Window) CellLayout() []container.Option {
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPerc(25, grid.Widget(w.Settings)),
			grid.ColWidthPercWithOpts(25, w.Cells[0].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[1].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[2].Layout()),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, w.Cells[3].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[4].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[5].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[6].Layout()),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, w.Cells[7].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[8].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[9].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[10].Layout()),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, w.Cells[11].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[12].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[13].Layout()),
			grid.ColWidthPercWithOpts(25, w.Cells[14].Layout()),
		),
	)
	opts, _ := builder.Build()
	return opts
}
