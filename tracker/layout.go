package tracker

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
)

type LayoutFunc func() []container.Option

func (t *Tracker) ChartLayout() []container.Option {
	layout := []container.Option{}
	layout = append(layout, t.MainOpts()...)
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(
						container.PlaceWidget(t.Settings),
					),
					container.Right(t.MetricScroller.Opts()...),
					container.SplitPercent(25),
				),
			),
			container.Bottom(
				container.PlaceWidget(t.MetricScroller.Current().Chart),
			),
			container.SplitPercent(30),
		),
	}
	layout = append(layout, opts...)
	return layout
}
func (t *Tracker) CellLayout() []container.Option {
	layout := []container.Option{}
	layout = append(layout, t.MainOpts()...)
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPerc(25, grid.Widget(t.Settings)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(0)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(1)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(2)),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(3)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(4)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(5)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(6)),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(7)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(8)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(9)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(10)),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(11)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(12)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(13)),
			grid.ColWidthPercWithOpts(25, t.GetMetricCell(14)),
		),
	)
	opts, _ := builder.Build()
	layout = append(layout, opts...)
	return layout
}
func (t *Tracker) SettingsLayout() []container.Option {
	layout := []container.Option{}
	layout = append(layout, t.MainOpts()...)
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPerc(50, grid.Widget(t.Chart)),
			grid.ColWidthPerc(50, grid.Widget(t.Cell)),
		),
	)
	builder.Add(
		grid.RowHeightPercWithOpts(25, t.ColorScroller.Opts()),
	)
	opts, _ := builder.Build()
	layout = append(layout, opts...)
	return layout
}
func (t *Tracker) MainOpts() []container.Option {
	return []container.Option{
		container.ID(t.name),
		container.Border(linestyle.Round),
		container.BorderTitle(" " + t.Name() + " "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(t.Color()),
		container.FocusedColor(t.Color()),
	}
}
func (t *Tracker) GetMetricCell(index int) []container.Option {
	if index >= len(t.Metrics) {
		return []container.Option{}
	}
	return t.Metrics[index].DisplayOpts()
}
