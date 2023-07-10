package tracker

// This file contains the implementation of Tracker layouts.

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
)

type LayoutFunc func() []container.Option

// initOpts creates opts for Tracker container.
func (t *Tracker) initOpts() []container.Option {
	return []container.Option{
		container.ID(t.name),
		container.Border(linestyle.Round),
		container.BorderTitle(" " + t.Name() + " "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(t.Color()),
		container.FocusedColor(t.Color()),
	}
}

// settingsLayout contains tickScroller, colorScroller, cell and chart buttons.
func (t *Tracker) settingsLayout() []container.Option {
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPerc(50, grid.Widget(t.chart)),
			grid.ColWidthPerc(50, grid.Widget(t.cell)),
		),
	)
	builder.Add(
		grid.RowHeightPercWithOpts(25, t.tick.Layout()),
		grid.RowHeightPercWithOpts(25, t.color.Layout()),
	)
	opts, _ := builder.Build()
	return opts
}

// focusLayout contains settings button, metricScroller and chart.
func (t *Tracker) focusLayout() []container.Option {
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(
						container.PlaceWidget(t.settings),
					),
					container.Right(t.metric.Layout()...),
					container.SplitPercent(25),
				),
			),
			container.Bottom(
				t.metric.Current().ChartLayout()...,
			),
			container.SplitPercent(30),
		),
	}
	return opts
}

// cellLayout contains up to 15 metrics and settings button.
func (t *Tracker) cellLayout() []container.Option {
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPerc(25, grid.Widget(t.settings)),
			grid.ColWidthPercWithOpts(25, t.metricCell(0)),
			grid.ColWidthPercWithOpts(25, t.metricCell(1)),
			grid.ColWidthPercWithOpts(25, t.metricCell(2)),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, t.metricCell(3)),
			grid.ColWidthPercWithOpts(25, t.metricCell(4)),
			grid.ColWidthPercWithOpts(25, t.metricCell(5)),
			grid.ColWidthPercWithOpts(25, t.metricCell(6)),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, t.metricCell(7)),
			grid.ColWidthPercWithOpts(25, t.metricCell(8)),
			grid.ColWidthPercWithOpts(25, t.metricCell(9)),
			grid.ColWidthPercWithOpts(25, t.metricCell(10)),
		),
	)
	builder.Add(
		grid.RowHeightPerc(25,
			grid.ColWidthPercWithOpts(25, t.metricCell(11)),
			grid.ColWidthPercWithOpts(25, t.metricCell(12)),
			grid.ColWidthPercWithOpts(25, t.metricCell(13)),
			grid.ColWidthPercWithOpts(25, t.metricCell(14)),
		),
	)
	opts, _ := builder.Build()
	return opts
}

// metricCell returns metric opts based on index in Tracker data.
//
// If index is not valid, returns empty opts.
func (t *Tracker) metricCell(index int) []container.Option {
	if index >= len(t.data) || index < 0 {
		return []container.Option{}
	}
	return t.data[index].DisplayLayout()
}
