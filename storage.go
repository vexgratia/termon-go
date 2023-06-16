package termon

import (
	"fmt"
	"runtime/metrics"
	"time"
)

type Storage struct {
	Tick    time.Duration
	Metrics map[string]*Metric
}

func NewStorage(tick time.Duration) *Storage {
	metrics := make(map[string]*Metric)
	for _, name := range AllMetrics {
		metrics[name] = NewMetric()
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
			metric.Data.Enqueue(float64(value.Uint64()))
		case metrics.KindFloat64:
			metric.Data.Enqueue(value.Float64())
		case metrics.KindFloat64Histogram:
			metric.Data.Enqueue(medianBucket(value.Float64Histogram()))
		case metrics.KindBad:
			panic("bug in runtime/metrics package!")
		default:
			fmt.Printf("%s: unexpected metric Kind: %v\n", name, value.Kind())
		}
		if metric.Data.Len() > metric.Capacity {
			metric.Data.Dequeue()
		}
	}
}
func (s *Storage) GetUpdates() {
	samples := make([]metrics.Sample, len(AllMetrics))
	for i, name := range AllMetrics {
		samples[i].Name = name
	}
	for {
		s.Update(&samples)
		time.Sleep(s.Tick)
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
