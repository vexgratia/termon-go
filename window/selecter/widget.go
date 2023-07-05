package selecter

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/widgets/button"
)

func (s *Selecter) MakeButtons() []*button.Button {
	buttons := []*button.Button{}
	for _, w := range s.Windows {
		b, _ := button.New(
			w.Name(),
			s.WindowSetFunc(w),
			button.Height(2),
			button.Width(8),
			button.FillColor(w.Color()),
		)
		buttons = append(buttons, b)
	}
	return buttons
}
func (s *Selecter) MakeContainers() [][]container.Option {
	containers := [][]container.Option{}
	for _, b := range s.Buttons {
		containers = append(containers, []container.Option{container.PlaceWidget(b)})
	}
	for len(containers) < 16 {
		containers = append(containers, []container.Option{})
	}
	return containers
}
