package selecter

// This file contains the implementation of Selecter and its basic methods.

import (
	"strconv"
	"sync"

	"github.com/mum4k/termdash/cell"
	"github.com/vexgratia/termon-go/updater"
	"github.com/vexgratia/termon-go/window"
)

// A Selecter is a Window type that allows to switch between other Windows.
type Selecter struct {
	// general
	id int
	// sync
	mu *sync.Mutex
	// data
	windows []window.Window
	// external
	updater *updater.Updater
}

// New creates a Selecter based on Updater.
func New(id int, updater *updater.Updater) *Selecter {
	return &Selecter{
		id:      id,
		mu:      &sync.Mutex{},
		updater: updater,
	}
}

// Add appends Window to Selecter.
func (s *Selecter) Add(w window.Window) {
	s.windows = append(s.windows, w)
}

// Name returns Selecter name.
func (s *Selecter) Name() string {
	return strconv.Itoa(s.id)
}

// Color returns Selecter color.
func (s *Selecter) Color() cell.Color {
	return cell.ColorWhite
}
