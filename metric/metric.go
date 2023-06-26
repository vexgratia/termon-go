package metric

import (
	"math"

	queue "github.com/vexgratia/collection-go/generic/queue"
)

const defaultCapacity = 100

type Metric struct {
	Name     string
	Current  float64
	Capacity uint32
	Queue    *queue.Queue[float64]
}

func New(name string) *Metric {
	queue := queue.New[float64]()
	for i := 0; i < defaultCapacity; i++ {
		queue.Enqueue(math.NaN())
	}
	return &Metric{
		Name:     name,
		Capacity: defaultCapacity,
		Queue:    queue,
		Current:  math.NaN(),
	}
}

func (m *Metric) Cap() int {
	return int(m.Capacity)
}

func (m *Metric) Resize(cap uint32) {
	if cap > m.Capacity {
		for i := m.Capacity; i <= cap; i++ {
			m.Queue.Enqueue(math.NaN())
		}
		m.Current = math.NaN()
	} else {
		for i := m.Capacity; i >= cap; i-- {
			m.Queue.Dequeue()
		}
	}
}
