package format

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
)

type FormatterFunc[T any] func(data T) []Text

type Text struct {
	Text string
	Opts []text.WriteOption
}

func Default[T any](data T) []Text {
	return []Text{
		{
			Text: fmt.Sprintf("%v", data),
			Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
		},
	}
}
