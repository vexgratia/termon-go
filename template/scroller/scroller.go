package scroller

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/text"
	dcl "github.com/vexgratia/collection-go/generic/list/dcl"
	"github.com/vexgratia/termon-go/format"
)

type Scroller[T any] struct {
	//
	name  string
	Color cell.Color
	//
	List *dcl.List[T]
	//
	Prev    *button.Button
	Display *text.Text
	Next    *button.Button
	//
	Formatter func(data T) []format.Text
	OnScroll  func()
}

func New[T any]() *Scroller[T] {
	scroller := &Scroller[T]{
		List:      dcl.New[T](),
		Formatter: format.Default[T],
		OnScroll:  func() {},
	}
	scroller.SetColor(cell.ColorWhite)
	return scroller
}
func (s *Scroller[T]) Name() string {
	return s.name
}
func (s *Scroller[T]) Add(data ...T) {
	for _, item := range data {
		s.List.Push(item)
	}
}
func (s *Scroller[T]) Current() T {
	return s.List.Peek()
}
func (s *Scroller[T]) ScrollNext() {
	s.List.ScrollNext()
}
func (s *Scroller[T]) ScrollPrev() {
	s.List.ScrollPrev()
}
func (s *Scroller[T]) SetScrollFunc(scroll func()) {
	s.OnScroll = scroll
}
func (s *Scroller[T]) SetFormatter(formatter func(data T) []format.Text) {
	s.Formatter = formatter
}
func (s *Scroller[T]) SetColor(color cell.Color) {
	s.Color = color
	s.ResetWidgets()
}
