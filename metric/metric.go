package metric

import (
	"math"

	queue "github.com/vexgratia/collection-go/generic/queue"
)

const maxCap = 10000

type Metric struct {
	Name     string
	Current  float64
	Capacity uint32
	Queue    *queue.Queue[float64]
}

func New(name string) *Metric {
	queue := queue.New[float64]()
	for i := 0; i < maxCap; i++ {
		queue.Enqueue(math.NaN())
	}
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
