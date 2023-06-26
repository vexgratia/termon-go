package cache

import (
	"fmt"
	"runtime/metrics"
	"time"
)

func (c *Cache) Update() {
	metrics.Read(c.Samples)
	for _, sample := range c.Samples {
		name, value := sample.Name, sample.Value
		metric := c.Metrics[name]
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
		metric.Resize(metric.Capacity)
	}
}
func (c *Cache) GetUpdates() {
	for {
		c.Update()
		time.Sleep(c.Tick)
	}
}
