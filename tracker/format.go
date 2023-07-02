package tracker

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/format"
	"github.com/vexgratia/termon-go/metric"
)

func MetricFormatter(m *metric.Metric) []format.TextWithOpts {
	chunk := []format.TextWithOpts{}
	tag := format.TextWithOpts{
		Text: fmt.Sprintf(" %s:\n", m.Tag),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	value := format.TextWithOpts{
		Text: m.CurrentF(),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(m.Color))},
	}
	chunk = append(chunk, tag, value)
	return chunk
}
