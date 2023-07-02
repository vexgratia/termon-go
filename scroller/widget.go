package scroller

import (
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/text"
)

func (s *Scroller[T]) ResetWidgets() {
	s.Prev, s.Next = s.MakeScrollButtons()
	s.Display = s.MakeDisplay()
}

func (s *Scroller[T]) MakeScrollButtons() (*button.Button, *button.Button) {
	prev, _ := button.New(
		"\u25C0\u2015\u2015\u2015",
		func() error {
			s.List.ScrollNext()
			s.Update()
			s.OnScroll()
			return nil
		},
		button.Height(2),
		button.FillColor(s.Color),
	)
	next, _ := button.New(
		"\u2015\u2015\u2015\u25B6",
		func() error {
			s.List.ScrollPrev()
			s.Update()
			s.OnScroll()
			return nil
		},
		button.Height(2),
		button.FillColor(s.Color),
	)
	return prev, next
}

func (s *Scroller[T]) MakeDisplay() *text.Text {
	display, _ := text.New(
		text.WrapAtWords(),
		text.DisableScrolling(),
	)
	return display
}
