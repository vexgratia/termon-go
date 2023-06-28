package cache

import (
	"fmt"
	"runtime/metrics"
)

func (c *Cache) Update() {
	var current float64
	metrics.Read(c.Samples)
	for _, sample := range c.Samples {
		name, value := sample.Name, sample.Value
		metric := c.Metrics[name]
		switch value.Kind() {
		case metrics.KindUint64:
			current = float64(value.Uint64())
		case metrics.KindFloat64:
			current = value.Float64()
		case metrics.KindFloat64Histogram:
			current = medianBucket(value.Float64Histogram())
		case metrics.KindBad:
			panic("bug in runtime/metrics package!")
		default:
			fmt.Printf("%s: unexpected metric Kind: %v\n", name, value.Kind())
		}
		metric.Add(current)
	}
}
