package scroller

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go"
)

func MetricDisplayFormatter(metric *termon.Metric) []FormatPair {
	pairs := []FormatPair{}
	pairs = append(pairs, FormatPair{
		Data: fmt.Sprintf("%s: ", metric.Name),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	})
	value := metric.Queue.Data[len(metric.Queue.Data)-1]
	pairs = append(pairs, FormatPair{
		Data: fmt.Sprintf("%f\n", value),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	})
	return pairs
}
