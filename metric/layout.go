package metric

// This file contains the implementation of Metric layouts.
import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

// DisplayLayout contains Metric display.
func (m *Metric) DisplayLayout() []container.Option {
	return []container.Option{
		container.ID(m.name),
		container.Border(linestyle.Light),
		container.BorderTitle(" " + m.tag + " "),
		container.BorderTitleAlignCenter(),
		container.BorderColor(m.color),
		container.FocusedColor(m.color),
		container.PlaceWidget(m.display),
	}
}

// ChartLayout contains Metric chart.
func (m *Metric) ChartLayout() []container.Option {
	return []container.Option{
		container.ID(m.name),
		container.Border(linestyle.Light),
		container.BorderTitle(" " + m.tag + " "),
		container.BorderTitleAlignCenter(),
		container.BorderColor(m.color),
		container.FocusedColor(m.color),
		container.PlaceWidget(m.chart),
	}
}
