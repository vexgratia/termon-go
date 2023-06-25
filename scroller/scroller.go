package scroller

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/text"
	dcl "github.com/vexgratia/collection-go/generic/list/dcl"
)

type Scroller[T any] struct {
	Color     cell.Color
	List      *dcl.List[T]
	Prev      *button.Button
	Display   *text.Text
	Next      *button.Button
	Formatter func(data T) []FormatPair
}
type FormatPair struct {
	Data string
	Opts []text.WriteOption
}

func MakeScroller[T any](data []T, color cell.Color, formatter func(data T) []FormatPair) *Scroller[T] {
	scroller := &Scroller[T]{
		Color:     color,
		Formatter: formatter,
	}
	list := dcl.New[T]()
	for _, data := range data {
		list.Push(data)
	}
	scroller.Prev, scroller.Next = scroller.MakeScrollButtons()
	scroller.Display = scroller.MakeDisplay()
	return scroller
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
func (s *Scroller[T]) Update() {
	pairs := s.Formatter(s.Current())
	s.Display.Reset()
	for _, pair := range pairs {
		s.Display.Write(pair.Data, pair.Opts...)
	}
}
