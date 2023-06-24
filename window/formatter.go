package window

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	. "github.com/vexgratia/termon-go/scroller"
)

var ModeName = map[WindowMode]string{

}

func ModeFormatter(mode WindowMode) []FormatPair {
	pairs := []FormatPair{}
	pairs = append(pairs, FormatPair{
		Data: fmt.Sprintf("%v", mode),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	})
	return pairs
}
