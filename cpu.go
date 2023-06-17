package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type CPU struct {
	Opts *[]container.Option
}

func NewCPU() *CPU {
	return &CPU{
		Opts: &[]container.Option{container.Border(linestyle.Round),
			container.BorderColor(cell.ColorRed),
			container.FocusedColor(cell.ColorRed),
			container.BorderTitle("CPU"),
			container.ID("cpu"),
			container.BorderTitleAlignCenter(),
		},
	}
}
