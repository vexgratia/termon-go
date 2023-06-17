package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type GC struct {
	Opts *[]container.Option
}

func NewGC() *GC {
	return &GC{
		Opts: &[]container.Option{container.Border(linestyle.Round),
			container.BorderColor(cell.ColorGreen),
			container.FocusedColor(cell.ColorGreen),
			container.BorderTitle("GC"),
			container.ID("gc"),
			container.BorderTitleAlignCenter(),
		},
	}
}
