package scroller

// This file contains the implementation of Scroller and its basic methods.

import (
	"sync"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/text"
	dcl "github.com/vexgratia/collection-go/generic/list/dcl"
	"github.com/vexgratia/termon-go/format"
)

type ScrollFunc func()

// A Scroller is a Template type that allows scrolling data of type T.
type Scroller[T any] struct {
	// general
	color cell.Color
	list  *dcl.List[T]
	// sync
	mu *sync.Mutex
	// widgets
	prev    *button.Button
	display *text.Text
	next    *button.Button
	// options
	formatter format.FormatterFunc[T]
	scroll    ScrollFunc
}

// New creates a Scroller of type T.
func New[T any]() *Scroller[T] {
	// general, sync and opts
	s := &Scroller[T]{
		color:     cell.ColorWhite,
		list:      dcl.New[T](),
		mu:        &sync.Mutex{},
		formatter: format.Default[T],
		scroll:    func() {},
	}
	// widgets
	s.reset()
	return s
}

// Add appends data to Scroller.
func (s *Scroller[T]) Add(data ...T) {
	for _, item := range data {
		s.list.Push(item)
		s.Update()
	}
}

// Color returns current Scroller color.
func (s *Scroller[T]) Color() cell.Color {
	return s.color
}

// Current returns Scroller current value.
func (s *Scroller[T]) Current() T {
	return s.list.Peek()
}

// ScrollNext moves head and tail nodes towards next.
func (s *Scroller[T]) scrollNext() {
	s.list.ScrollNext()
}

// ScrollPrev moves head and tail nodes towards previous.
func (s *Scroller[T]) scrollPrev() {
	s.list.ScrollPrev()
}
