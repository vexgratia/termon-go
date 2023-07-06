package tracker

// This file contains the implementation of Tracker internal formatting features.

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/format"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/palette"
)

// metricFormat returns formatted text for Scroller display.
func metricFormat(m *metric.Metric) []format.Text {
	chunk := []format.Text{}
	tag := format.Text{
		Text: fmt.Sprintf(" %s:\n", m.Tag()),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	value := format.Text{
		Text: m.CurrentF(),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(m.Color()))},
	}
	chunk = append(chunk, tag, value)
	return chunk
}

// colorFormat returns formatted text for Scroller display.
func colorFormat(c cell.Color) []format.Text {
	chunk := []format.Text{}
	tag := format.Text{
		Text: fmt.Sprintf(" Color: "),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	value := format.Text{
		Text: " " + palette.String(c),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(c))},
	}
	chunk = append(chunk, tag, value)
	return chunk
}
