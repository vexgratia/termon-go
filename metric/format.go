package metric

import "fmt"

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

func (m *Metric) Tag() string {
	return m.Parsed.Tag
}

func (m *Metric) CurrentF() string {
	switch m.Parsed.Type {
	case SECONDS:
		return FormatSeconds(m.Current)
	case BYTES:
		return FormatBytes(m.Current)
	case CALLS:
		return FormatCalls(m.Current)
	case OBJECTS:
		return FormatObjects(m.Current)
	case NUMBER:
		return FormatNumber(m.Current)
	case CYCLE:
		return FormatCycle(m.Current)
	case CYCLES:
		return FormatCycles(m.Current)
	case THREADS:
		return FormatThreads(m.Current)
	}
	return fmt.Sprintf("%v", m.Current)
}

func FormatBytes(value float64) string {
	var format string
	if value < 1024 {
		return fmt.Sprintf(" %5.1f", value) + " B"
	}
	value /= 1024
	for _, char := range "KMGTPEZY" {
		if value < 1024 {
			format += " " + string(char)
			break
		}
		value /= 1024
	}
	format = fmt.Sprintf(" %5.1f", value) + format + "B"
	return format
}
func FormatSeconds(value float64) string {
	var suffix string
	if value > 84600 {
		value /= 84600
		suffix = "D"
	} else if value > 3600 {
		value /= 3600
		suffix = "H"
	} else if value > 60 {
		value /= 60
		suffix = "M"
	} else if value > 1 {
		suffix = "S"
	} else if value > 0.001 {
		value *= 1000
		suffix = "MS"
	} else if value > 0.000001 {
		value *= 1000000
		suffix = "mS"
	} else {
		value *= 1000000000
		suffix = "nS"
	}
	return fmt.Sprintf(" %5.1f %s", value, suffix)
}
func FormatCalls(value float64) string {
	return fmt.Sprintf(" %5.0G calls", value)
}
func FormatObjects(value float64) string {
	return fmt.Sprintf(" %5.0G objects", value)
}
func FormatNumber(value float64) string {
	return fmt.Sprintf(" %5.0f", value)
}
func FormatCycle(value float64) string {
	return fmt.Sprintf("%5.0f cycle", value)
}
func FormatCycles(value float64) string {
	return fmt.Sprintf(" %5.0G cycles", value)
}
func FormatThreads(value float64) string {
	return fmt.Sprintf(" %5.0f threads", value)
}
