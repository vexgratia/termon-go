package logger

import (
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"
)

func (l *Logger) ResetWidgets() {
	l.Spark = l.MakeSpark()
	l.Display = l.MakeDisplay()
}

func (l *Logger) MakeSpark() *sparkline.SparkLine {
	spark, _ := sparkline.New(
		sparkline.Label(" Tick: "),
		sparkline.Color(l.color),
	)
	return spark
}
func (l *Logger) MakeDisplay() *text.Text {
	text, _ := text.New(
		text.WrapAtWords(),
	)
	return text
}
