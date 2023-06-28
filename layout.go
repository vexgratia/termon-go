package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type LayoutFunc func() []container.Option

func (t *Termon) DefaultLayout() []container.Option {
	return []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(t.CPU.Opts()...),
					container.Right(t.GC.Opts()...),
				),
			),
			container.Bottom(
				container.SplitVertical(
					container.Left(t.Golang.Opts()...),
					container.Right(t.Memory.Opts()...),
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
