package termon

import (
	"fmt"
	"runtime/metrics"
	"time"

	metric "github.com/vexgratia/termon-go/metric"
)

type Storage struct {
	Tick    time.Duration
	Metrics map[string]*metric.Metric
}

func New(tick time.Duration, names []string) *Storage {
	metrics := make(map[string]*metric.Metric)
	for _, name := range names {
		metrics[name] = metric.New()
	}
	return &Storage{
		Tick:    tick,
		Metrics: metrics,
	}
}
func (s *Storage) Update(samples *[]metrics.Sample) {
	metrics.Read(*samples)
	for _, sample := range *samples {
		name, value := sample.Name, sample.Value
		metric := s.Metrics[name]
		switch value.Kind() {
		case metrics.KindUint64:
			metric.Queue.Enqueue(float64(value.Uint64()))
		case metrics.KindFloat64:
			metric.Queue.Enqueue(value.Float64())
		case metrics.KindFloat64Histogram:
			metric.Queue.Enqueue(medianBucket(value.Float64Histogram()))
		case metrics.KindBad:
			panic("bug in runtime/metrics package!")
		default:
			fmt.Printf("%s: unexpected metric Kind: %v\n", name, value.Kind())
		}
		if metric.Queue.Len() > metric.Capacity {
			metric.Queue.Dequeue()
		}
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