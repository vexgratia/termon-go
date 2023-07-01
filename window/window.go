package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
)

type Window interface {
	Name() string
	Color() cell.Color
	Lock()
	Unlock()
	Opts() []container.Option
	Run()
}