package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type Memory struct {
	Opts *[]container.Option
}

func NewMemory() *Memory {
	return &Memory{
		Opts: &[]container.Option{container.Border(linestyle.Round),
			container.BorderColor(cell.ColorYellow),
			container.FocusedColor(cell.ColorYellow),
			container.BorderTitle("Memory"),
			container.ID("memory"),
			container.BorderTitleAlignCenter(),
		},
	}
}
