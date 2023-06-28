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
	Current    float64
	Sum        float64
	Max        float64
	Avg        *queue.Queue[float64]
	CurrentAvg float64
	//
	Capacity uint32
	Data     *queue.Queue[float64]
	//
	Format func(value float64) string
}

func New(name string) *Metric {
	data := queue.New[float64]()
	avg := queue.New[float64]()
	return &Metric{
		Name:     name,
		Capacity: maxCap,
		Avg:      avg,
		Data:     data,
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
	m.Current = value
	m.Sum += value
	m.Data.Enqueue(m.Current)
	for m.Data.Len() > m.Cap() {
		deq := m.Data.Dequeue()
		m.Sum -= deq
		m.Avg.Dequeue()
	}
	m.CurrentAvg = m.Sum / float64(m.Cap())
	m.Avg.Enqueue(m.CurrentAvg)
}
