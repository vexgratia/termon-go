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

func New(tick time.Duration) *Cache {
	samples := make([]metrics.Sample, len(metric.All))
	metrics := make(map[string]*metric.Metric)
	for i, name := range metric.All {
		metrics[name] = metric.New(name)
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
func (c *Cache) GetMetrics(names []string) []*metric.Metric {
	metrics := []*metric.Metric{}
	for _, name := range names {
		metrics = append(metrics, c.Metrics[name])
	}
	return metrics
}
