package window

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/format"
	"github.com/vexgratia/termon-go/metric"
)

func (w *Window) MetricFormat(m *metric.Metric) []format.TextWithOpts {
	metric := format.TextWithOpts{
		Text: m.Tag() + ":\n",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	value := format.TextWithOpts{
		Text: m.CurrentF(),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(w.Color))},
	}
	return []format.TextWithOpts{metric, value}
}
func (w *Window) CapFormat(c uint32) []format.TextWithOpts {
	init := format.TextWithOpts{
		Text: "\nDisplay capacity:\n",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	cap := format.TextWithOpts{
		Text: fmt.Sprintf("%d ", c),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(w.Color))},
	}
	ticks := format.TextWithOpts{
		Text: "ticks",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	return []format.TextWithOpts{init, cap, ticks}
}
