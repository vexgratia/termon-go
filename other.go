package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type Other struct {
	Opts *[]container.Option
}

func NewOther() *Other {
	return &Other{
		Opts: &[]container.Option{container.Border(linestyle.Round),
			container.BorderColor(cell.ColorBlue),
			container.FocusedColor(cell.ColorBlue),
			container.BorderTitle("Other"),
			container.ID("other"),
			container.BorderTitleAlignCenter(),
		},
	}
}
