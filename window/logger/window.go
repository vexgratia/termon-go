package logger

import (
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
)

func (l *Logger) Name() string {
	return l.name
}
func (l *Logger) Color() cell.Color {
	return l.color
}
func (l *Logger) Opts() []container.Option {
	return l.Layout()
}
func (l *Logger) GetUpdates() {
	l.Add("Start logging...")
	for {
		l.Update()
		time.Sleep(time.Millisecond * 10)
	}
}
