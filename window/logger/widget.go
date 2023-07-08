package logger

// This file contains the implementation of Logger widgets.

import (
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"
)

// reset recreates all Logger widgets.
//
// Blocks Logger mutex to avoid data race.
func (l *Logger) reset() {
	l.spark = l.makeSpark()
	l.display = l.makeDisplay()
}

// makeSpark creates Logger spark.
func (l *Logger) makeSpark() *sparkline.SparkLine {
	spark, _ := sparkline.New(
		sparkline.Label(" Tick: "),
		sparkline.Color(l.Color()),
	)
	return spark
}

// makeDisplay creates Logger display.
func (l *Logger) makeDisplay() *text.Text {
	text, _ := text.New(
		text.WrapAtWords(),
		text.RollContent(),
	)
	return text
}
