package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
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
					container.Right(
						w.MetricScroller.Layout()...,
					),
					container.SplitPercent(20),
				),
			),
			container.Bottom(
				container.PlaceWidget(w.ModeWidget()),
			),
			container.SplitPercent(30),
		),
	}
}
func (w *Window) SettingsLayout() []container.Option {
	return []container.Option{
		container.ID(w.Name),
		container.Border(linestyle.Round),
		container.BorderTitle(w.Name + " SETTINGS"),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorGray),
		container.FocusedColor(cell.ColorGray),
		container.SplitHorizontal(
			container.Top(
				container.PlaceWidget(w.Return),
			),
			container.Bottom(
				w.ModeScroller.Layout()...,
			),
		),
	}
}
