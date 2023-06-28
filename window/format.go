package window

import (
	"fmt"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/format"
	"github.com/vexgratia/termon-go/metric"
)

func (w *Window) MetricScrollerFormat(m *metric.Metric) []format.TextWithOpts {
	result := []format.TextWithOpts{}
	//
	metric := format.TextWithOpts{
		Text: " " + m.Tag() + ": ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	value := format.TextWithOpts{
		Text: m.Format(m.Current),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(w.Color))},
	}
	result = append(result, metric, value)
	//
	avg := format.TextWithOpts{
		Text: "\n Average: ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	avgValue := format.TextWithOpts{
		Text: m.Format(m.CurrentAvg),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(w.Color))},
	}
	if w.ShowAvg.Current() {
		result = append(result, avg, avgValue)
	}
	//
	max := format.TextWithOpts{
		Text: "\n Max: ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	maxValue := format.TextWithOpts{
		Text: m.Format(m.Max),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(w.Color))},
	}
	if w.ShowMax.Current() {
		result = append(result, max, maxValue)
	}
	return result
}
func (w *Window) CapScrollerFormat(c uint32) []format.TextWithOpts {
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
func (w *Window) TickScrollerFormat(t time.Duration) []format.TextWithOpts {
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
func (w *Window) AvgScrollerFormat(b bool) []format.TextWithOpts {
	init := format.TextWithOpts{
		Text: " Show average: ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	show := format.TextWithOpts{
		Text: fmt.Sprintf("%v ", b),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(w.Color))},
	}
	return []format.TextWithOpts{init, show}
}
func (w *Window) MaxScrollerFormat(b bool) []format.TextWithOpts {
	init := format.TextWithOpts{
		Text: " Show maximum: ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	show := format.TextWithOpts{
		Text: fmt.Sprintf("%v ", b),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(w.Color))},
	}
	return []format.TextWithOpts{init, show}
}
