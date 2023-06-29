package tracker

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type LayoutFunc func() []container.Option

func (t *Tracker) Opts() []container.Option {
	return t.Layout()
}
func (t *Tracker) ChartLayout() []container.Option {
	return []container.Option{
		container.ID(t.name),
		container.Border(linestyle.Round),
		container.BorderTitle(" " + t.Name() + " "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(t.Color),
		container.FocusedColor(t.Color),
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(
						container.PlaceWidget(t.Settings),
					),
					container.Right(t.MetricScroller.Opts()...),
					container.SplitPercent(20),
				),
			),
			container.Bottom(
				container.PlaceWidget(t.MetricScroller.Current().Chart),
			),
			container.SplitPercent(30),
		),
	}
}
