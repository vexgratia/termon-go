package window

import (
	"fmt"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/format"
	"github.com/vexgratia/termon-go/metric"
)

func (w *Window) MetricFormat(m *metric.Metric) []format.TextWithOpts {
	metric := format.TextWithOpts{
		Text: " " + m.Tag() + ": ",
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
		Text: " Display capacity: ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	cap := format.TextWithOpts{
		Text: fmt.Sprintf("%d ", c),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(w.Color))},
	}
	ticks := format.TextWithOpts{
		Text: "tick",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	return []format.TextWithOpts{init, cap, ticks}
}
func (w *Window) TickFormat(t time.Duration) []format.TextWithOpts {
	init := format.TextWithOpts{
		Text: " Tick: ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	tick := format.TextWithOpts{
		Text: fmt.Sprintf("%d ", t),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(w.Color))},
	}
	ms := format.TextWithOpts{
		Text: "mS",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	return []format.TextWithOpts{init, tick, ms}
}
