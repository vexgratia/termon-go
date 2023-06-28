package metric

import (
	"math"

	queue "github.com/vexgratia/collection-go/generic/queue"
)

const maxCap = 10000

type Metric struct {
	Name   string
	Parsed Parsed
	//
	Max     float64
	Current float64
	//
	Capacity uint32
	Queue    *queue.Queue[float64]
}

func New(name string) *Metric {
	queue := queue.New[float64]()
	return &Metric{
		Name:     name,
		Capacity: maxCap,
		Queue:    queue,
		Current:  math.NaN(),
	}
}

func (m *Metric) Cap() int {
	return int(m.Capacity)
}
func (m *Metric) Add(value float64) {
	if value > m.Max {
		m.Max = value
	}
	m.Queue.Enqueue(value)
	for m.Queue.Len() > m.Cap() {
		m.Queue.Dequeue()
	}
	m.Current = value
}
