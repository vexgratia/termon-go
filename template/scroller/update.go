package scroller

import (
	"github.com/mum4k/termdash/cell"
	"github.com/vexgratia/termon-go/format"
)

// Update updates Scroller widgets.
func (s *Scroller[T]) Update() {
	s.display.Reset()
	textChunk := s.formatter(s.Current())
	for _, text := range textChunk {
		s.display.Write(text.Text, text.Opts...)
	}
}

// SetScrollFunc sets Scroller scroll to given ScrollFunc.
//
// Calls reset.
func (s *Scroller[T]) SetScrollFunc(fn ScrollFunc) {
	s.scroll = fn
	s.reset()
}

// SetColor sets Scroller color to given color.
//
// Calls reset.
func (s *Scroller[T]) SetColor(color cell.Color) {
	s.color = color
	s.reset()
}

// SetFormatter sets Scroller formatter to given FormatterFunc.
//
// Calls reset.
func (s *Scroller[T]) SetFormatter(fn func(data T) []format.Text) {
	s.formatter = fn
	s.reset()
}
