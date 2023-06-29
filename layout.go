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
	layout := []container.Option{}
	layout = append(layout, t.MainOpts()...)
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(),
			container.Bottom(t.Windows[0].Opts()...),
			container.SplitPercent(1),
		),
	}
	layout = append(layout, opts...)
	return layout
}
func (t *Termon) LayoutTwoHorizontal() []container.Option {
	layout := []container.Option{}
	layout = append(layout, t.MainOpts()...)
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(t.Windows[0].Opts()...),
			container.Bottom(t.Windows[1].Opts()...),
			container.SplitPercent(50),
		),
	}
	layout = append(layout, opts...)
	return layout
}
func (t *Termon) LayoutTwoVertical() []container.Option {
	layout := []container.Option{}
	layout = append(layout, t.MainOpts()...)
	opts := []container.Option{
		container.SplitVertical(
			container.Left(t.Windows[0].Opts()...),
			container.Right(t.Windows[1].Opts()...),
			container.SplitPercent(50),
		),
	}
	layout = append(layout, opts...)
	return layout
}
func (t *Termon) LayoutFour() []container.Option {
	layout := []container.Option{}
	layout = append(layout, t.MainOpts()...)
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(t.Windows[0].Opts()...),
					container.Right(t.Windows[1].Opts()...),
					container.SplitPercent(50),
				),
			),
			container.Bottom(
				container.SplitVertical(
					container.Left(t.Windows[2].Opts()...),
					container.Right(t.Windows[3].Opts()...),
					container.SplitPercent(50),
				),
			),
			container.SplitPercent(50),
		),
	}
	layout = append(layout, opts...)
	return layout
}
func (t *Termon) MakeMain() *container.Container {
	main, _ := container.New(
		t.Terminal,
		t.MainOpts()...,
	)
	return main
}
func (t *Termon) MainOpts() []container.Option {
	return []container.Option{
		container.ID("MAIN"),
		container.Border(linestyle.Round),
		container.BorderTitle(" TERMON "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
	}
}
