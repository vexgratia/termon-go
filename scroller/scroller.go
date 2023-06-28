package scroller

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/text"
	dcl "github.com/vexgratia/collection-go/generic/list/dcl"
	format "github.com/vexgratia/termon-go/format"
)

type Scroller[T any] struct {
	Name      string
	Color     cell.Color
	List      *dcl.List[T]
	Prev      *button.Button
	Display   *text.Text
	Next      *button.Button
	Formatter func(data T) []format.TextWithOpts
	Scroll    func()
}

func New[T any](data []T, color cell.Color, formatter func(data T) []format.TextWithOpts, scroll func()) *Scroller[T] {
	scroller := &Scroller[T]{
		Color:     color,
		Formatter: formatter,
		Scroll:    scroll,
	}
	list := dcl.New[T]()
	for _, data := range data {
		list.Push(data)
	}
	scroller.List = list
	scroller.Prev, scroller.Next = scroller.MakeScrollButtons()
	scroller.Display = scroller.MakeDisplay()
	scroller.Update()
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
