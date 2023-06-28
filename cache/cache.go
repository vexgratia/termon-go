package cache

import (
	"runtime/metrics"
	"time"

	metric "github.com/vexgratia/termon-go/metric"
)

type Cache struct {
	Tick    time.Duration
	Metrics map[string]*metric.Metric
	Samples []metrics.Sample
}

func New(names []string, tick time.Duration) *Cache {
	samples := make([]metrics.Sample, len(names))
	metrics := make(map[string]*metric.Metric)
	for i, name := range names {
		metrics[name] = metric.New(name)
		metrics[name].Parsed = metrics[name].Parse()
		samples[i].Name = name
	}
	return &Cache{
		Tick:    tick,
		Metrics: metrics,
		Samples: samples,
	}
}
func (c *Cache) SetTick(tick time.Duration) {
	c.Tick = tick
}
func (c *Cache) GetMetrics() []*metric.Metric {
	metrics := []*metric.Metric{}
	for _, metric := range c.Metrics {
		metrics = append(metrics, metric)
	}
	return metrics
}
