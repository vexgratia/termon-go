package cache

import (
	"runtime/metrics"

	metric "github.com/vexgratia/termon-go/metric"
)

type Cache struct {
	Metrics map[string]*metric.Metric
	Samples []metrics.Sample
}

func New(names []string) *Cache {
	samples := make([]metrics.Sample, len(names))
	metrics := make(map[string]*metric.Metric)
	for i, name := range names {
		metrics[name] = metric.New(name)
		metrics[name].Parsed = metrics[name].Parse()
		metrics[name].Format = metrics[name].Formatter()
		samples[i].Name = name
	}
	return &Cache{
		Metrics: metrics,
		Samples: samples,
	}
}

func (c *Cache) GetMetrics() []*metric.Metric {
	metrics := []*metric.Metric{}
	for _, metric := range c.Metrics {
		metrics = append(metrics, metric)
	}
	return metrics
}
