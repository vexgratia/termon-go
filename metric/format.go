package metric

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/format"
)

type MetricType int

const (
	SECONDS MetricType = iota
	BYTES
	CALLS
	OBJECTS
	NUMBER
	CYCLES
	CYCLE
	THREADS
)

func MakeValueFormatter(t MetricType) func(value float64) string {
	switch t {
	case SECONDS:
		return SecondsFormat
	case BYTES:
		return BytesFormat
	default:
		return NumFormat
	}
}
func (m *Metric) DisplayFormat() []format.TextWithOpts {
	return []format.TextWithOpts{
		{
			Text: m.CurrentF(),
			Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
		},
	}
}

func SecondsFormat(value float64) string {
	if value < 1 {
		return LessSecondFormat(value)
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
func LessSecondFormat(value float64) string {
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
func BytesFormat(value float64) string {
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
func NumFormat(value float64) string {
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
