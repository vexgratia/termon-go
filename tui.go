package termon

import (
	"fmt"
	"runtime/metrics"
	"time"

	. "github.com/vexgratia/collection-go/models/generic/queue"
)

type TUI struct {
	Tick     time.Duration
	Interval time.Duration
	Metrics  map[string]*Queue[float64]
}

func NewTUI(tick int, interval int) *TUI {
	return &TUI{
		Tick:     time.Duration(tick) * time.Millisecond * 150,
		Interval: time.Duration(interval) * time.Second,
		Metrics:  make(map[string]*Queue[float64]),
	}
}

func (t *TUI) GetUpdates() {
	descs := metrics.All()
	samples := make([]metrics.Sample, len(descs))
	for _, des := range descs {
		t.Metrics[des.Name] = NewQueue[float64]()
	}
	for {
		descs = metrics.All()
		for i := range samples {
			samples[i].Name = descs[i].Name
		}
		metrics.Read(samples)
		for _, sample := range samples {
			name, value := sample.Name, sample.Value
			queue := t.Metrics[name]
			switch value.Kind() {
			case metrics.KindUint64:
				queue.Enqueue(float64(value.Uint64()))
			case metrics.KindFloat64:
				queue.Enqueue(value.Float64())
			case metrics.KindFloat64Histogram:
				queue.Enqueue(medianBucket(value.Float64Histogram()))
			case metrics.KindBad:
				panic("bug in runtime/metrics package!")
			default:
				fmt.Printf("%s: unexpected metric Kind: %v\n", name, value.Kind())
			}
			if queue.Len() > int(t.Interval)/100000000 {
				queue.Dequeue()
			}
		}

		time.Sleep(t.Tick)
	}
}

func medianBucket(h *metrics.Float64Histogram) float64 {
	total := uint64(0)
	for _, count := range h.Counts {
		total += count
	}
	thresh := total / 2
	total = 0
	for i, count := range h.Counts {
		total += count
		if total >= thresh {
			return h.Buckets[i]
		}
	}
	panic("should not happen")
}
