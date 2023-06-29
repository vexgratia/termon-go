package metric

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

func (m *Metric) DisplayOpts() []container.Option {
	return []container.Option{
		container.ID(m.name),
		container.Border(linestyle.Light),
		container.BorderTitle(" " + m.Tag + " "),
		container.BorderTitleAlignCenter(),
		container.BorderColor(m.Color),
		container.FocusedColor(m.Color),
		container.PlaceWidget(m.Display),
	}
}
func (m *Metric) ChartOpts() []container.Option {
	return []container.Option{
		container.ID(m.name),
		container.Border(linestyle.Light),
		container.BorderTitle(" " + m.Tag + " "),
		container.BorderTitleAlignCenter(),
		container.BorderColor(m.Color),
		container.FocusedColor(m.Color),
		container.PlaceWidget(m.Chart),
	}
}
