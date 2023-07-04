package tracker

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/format"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/palette"
)

func MetricFormatter(m *metric.Metric) []format.Text {
	chunk := []format.Text{}
	tag := format.Text{
		Text: fmt.Sprintf(" %s:\n", m.Tag),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	value := format.Text{
		Text: m.CurrentF(),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(m.Color))},
	}
	chunk = append(chunk, tag, value)
	return chunk
}
func ColorFormatter(c cell.Color) []format.Text {
	chunk := []format.Text{}
	tag := format.Text{
		Text: fmt.Sprintf(" Color: \n"),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	value := format.Text{
		Text: " " + palette.String(c),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(c))},
	}
	chunk = append(chunk, tag, value)
	return chunk
}
