package logger

// This file contains the implementation of Logger templates.

import (
	"github.com/mum4k/termdash/cell"
	"github.com/vexgratia/termon-go/palette"
	"github.com/vexgratia/termon-go/template/scroller"
)

// makeColorScroller creates a Scroller for Tracker colors.
func (l *Logger) makeColorScroller() *scroller.Scroller[cell.Color] {
	color := scroller.New[cell.Color]()
	color.SetScrollFunc(l.reColor)
	color.SetFormatter(colorFormat)
	color.Add(palette.All...)
	return color
}
