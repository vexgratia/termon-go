package scroller

// This file contains the implementation of Scroller widgets.

import (
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/text"
)

// reset recreates all Scroller widgets.
//
// Blocks Scroller mutex to avoid data race.
func (s *Scroller[T]) reset() {
	s.mu.Lock()
	s.prev, s.next = s.makeScrollButtons()
	s.display = s.makeDisplay()
	s.mu.Unlock()
}

// makeScrollButtons creates Scroller next and prev buttons.
func (s *Scroller[T]) makeScrollButtons() (*button.Button, *button.Button) {
	prev, _ := button.New(
		"\u25C0\u2015\u2015\u2015",
		func() error {
			s.scrollNext()
			s.scroll()
			s.Update()
			return nil
		},
		button.Height(2),
		button.FillColor(s.Color()),
	)
	next, _ := button.New(
		"\u2015\u2015\u2015\u25B6",
		func() error {
			s.scrollPrev()
			s.scroll()
			s.Update()
			return nil
		},
		button.Height(2),
		button.FillColor(s.Color()),
	)
	return prev, next
}

// makeDisplay creates Scroller display.
func (s *Scroller[T]) makeDisplay() *text.Text {
	display, _ := text.New(
		text.WrapAtWords(),
		text.DisableScrolling(),
	)
	return display
}
