package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type LayoutFunc func() []container.Option

func (t *Termon) Opts() []container.Option {
	return t.Layout()
}

func (t *Termon) LayoutOne() []container.Option {
	return []container.Option{
		container.SplitHorizontal(
			container.Top(t.Windows[0].Opts()...),
			container.Bottom(),
			container.SplitPercent(90),
		),
	}
}
func (t *Termon) LayoutTwoHorizontal() []container.Option {
	return []container.Option{
		container.SplitHorizontal(
			container.Top(t.Windows[0].Opts()...),
			container.Bottom(t.Windows[1].Opts()...),
		),
	}
}
func (t *Termon) LayoutTwoVertical() []container.Option {
	return []container.Option{
		container.SplitVertical(
			container.Left(t.Windows[0].Opts()...),
			container.Right(t.Windows[1].Opts()...),
		),
	}
}
func (t *Termon) LayoutFour() []container.Option {
	return []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(t.Windows[0].Opts()...),
					container.Right(t.Windows[1].Opts()...),
				),
			),
			container.Bottom(
				container.SplitVertical(
					container.Left(t.Windows[2].Opts()...),
					container.Right(t.Windows[3].Opts()...),
				),
			),
		),
	}
}
func (t *Termon) MakeMain() *container.Container {
	main, _ := container.New(
		t.Terminal,
		container.ID("MAIN"),
		container.Border(linestyle.Round),
		container.BorderTitle(" TERMON "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
	)
	return main
}
