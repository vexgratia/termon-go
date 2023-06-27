package scroller

import (
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/text"
)

func (s *Scroller[T]) MakeScrollButtons() (*button.Button, *button.Button) {
	prev, _ := button.New(
		"<---",
		func() error {
			s.List.ScrollNext()
			s.Update()
			return nil
		},
		button.Height(2),
		button.FillColor(s.Color),
	)
	next, _ := button.New(
		"--->",
		func() error {
			s.List.ScrollPrev()
			s.Update()
			return nil
		},
		button.Height(2),
		button.FillColor(s.Color),
	)
	return prev, next
}

func (s *Scroller[T]) MakeDisplay() *text.Text {
	display, _ := text.New(
		text.WrapAtRunes(),
		text.DisableScrolling(),
	)
	return display
}
