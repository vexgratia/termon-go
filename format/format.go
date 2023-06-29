package format

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
)

type TextWithOpts struct {
	Text string
	Opts []text.WriteOption
}

func Default[T any](data T) []TextWithOpts {
	return []TextWithOpts{
		{
			Text: fmt.Sprintf("%v", data),
			Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
		},
	}
}
