package metric

import (
	"fmt"
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

func Formatter(t MetricType) func(m *Metric) string {
	switch t {
	// ...
	}
	return func(m *Metric) string {
		return fmt.Sprintf("%v", m.Current())
	}
}
