package metric

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/format"
)

type ValueFormatter func(value float64) string

type Unit int

const (
	SECONDS Unit = iota
	BYTES
	CALLS
	OBJECTS
	NUMBER
	CYCLES
	CYCLE
	THREADS
)

// displayFormat returns formatted text for Metric display.
func (m *Metric) displayFormat() []format.Text {
	return []format.Text{
		{
			Text: m.CurrentF(),
			Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
		},
	}
}

// MetricFormatter returns valid ValueFormatter based on Unit.
func MetricFormatter(t Unit) ValueFormatter {
	switch t {
	case SECONDS:
		return SecondsFormatter
	case BYTES:
		return BytesFormatter
	default:
		return NumFormatter
	}
}

// SecondsFormatter is a ValueFormatter for SECONDS.
func SecondsFormatter(value float64) string {
	if value < 1 {
		return LessSecondFormatter(value)
	}
	if value < 60 {
		return fmt.Sprintf("%6.1f s", value)
	}
	if value < 3600 {
		m := int(value) / 60
		s := value - float64(m*60)
		return fmt.Sprintf("%4d:%02.0f m", m, s)
	}
	h := int(value) / 3600
	m := (value - float64(h*3600)) / 60
	return fmt.Sprintf("%4d:%02.0f h", h, m)
}

// LessSecondFormatter is a helper function for SecondsFormatter.
func LessSecondFormatter(value float64) string {
	var format string
	value *= 1000000000
	for _, char := range "nµm" {
		if value < 1000 {
			format += " " + string(char)
			break
		}
		value /= 1000
	}
	return fmt.Sprintf("%6.1f", value) + format + "s"
}

// BytesFormatter is a ValueFormatter for BYTES.
func BytesFormatter(value float64) string {
	var format string
	if value < 1024 {
		return fmt.Sprintf("%6.1f", value) + " B"
	}
	value /= 1024
	for _, char := range "KMGTPEZY" {
		if value < 1024 {
			format += " " + string(char)
			break
		}
		value /= 1024
	}
	format = fmt.Sprintf("%6.1f", value) + format + "B"
	return format
}

// NumFormatter is a ValueFormatter for any integer Unit.
func NumFormatter(value float64) string {
	var format string
	if value < 1000 {
		return fmt.Sprintf("%7.0f", value)
	}
	value /= 1000
	for _, char := range "KMGTPE" {
		if value < 1000 {
			format += " " + string(char)
			break
		}
		value /= 1000
	}
	format = fmt.Sprintf("%6.1f", value) + format
	return format
}
