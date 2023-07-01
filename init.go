package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

func (t *Termon) MakeMain() *container.Container {
	main, _ := container.New(
		t.Terminal,
		container.ID("Termon"),
		container.Border(linestyle.Round),
		container.BorderTitle(" TERMON "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
		container.SplitHorizontal(
			container.Top(),
			container.Bottom(container.ID("Main")),
			container.SplitPercent(1),
		),
	)
	return main
}
