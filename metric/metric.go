package metric

// This file contains the implementation of Metric and its basic methods.

import (
	"math"
	"sync"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
	queue "github.com/vexgratia/collection-go/generic/queue"
)

var maxCap = 1000

// A Metric is a data structure to store and display Golang runtime data.
//
// Based on float64 Queue.
type Metric struct {
	// general
	name  string
	color cell.Color
	// parsed
	tag  string
	unit Unit
	// sync
	mu *sync.Mutex
	// data
	current float64
	data    *queue.Queue[float64]
	// widgets
	display *text.Text
	chart   *linechart.LineChart
	// options
	formatter ValueFormatter
}

// New creates a Metric based on given name.
func New(name string) *Metric {
	m := &Metric{
		// general
		name:  name,
		color: cell.ColorWhite,
		// sync
		mu: &sync.Mutex{},
		// data
		data: queue.New[float64](),
	}
	m.parse()
	for i := 0; i < maxCap; i++ {
		m.data.Enqueue(math.NaN())
	}
	m.reset()
	return m
}

// Name returns Metric name.
func (m *Metric) Name() string {
	return m.name
}

// Color returns Metric color.
func (m *Metric) Color() cell.Color {
	return m.color
}

// Tag returns Metric tag.
func (m *Metric) Tag() string {
	return m.name
}

// Unit returns Metric unit.
func (m *Metric) Unit() string {
	return m.name
}

// Current returns current metric value.
func (m *Metric) Current() float64 {
	return m.current
}

// CurrentF returns current metric value, formatted with metric formatter.
func (m *Metric) CurrentF() string {
	return m.formatter(m.current)
}
