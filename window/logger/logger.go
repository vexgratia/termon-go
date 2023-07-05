package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/collection-go/generic/queue"
	"github.com/vexgratia/termon-go/format"
	"github.com/vexgratia/termon-go/updater"
)

var maxCap = 200

type Logger struct {
	name     string
	color    cell.Color
	Layout   LayoutFunc
	Data     *queue.Queue[string]
	Settings *button.Button
	Display  *text.Text
	Spark    *sparkline.SparkLine
	Updater  *updater.Updater
}

func New(name string, updater *updater.Updater) *Logger {
	logger := &Logger{
		name:    name,
		Updater: updater,
		Data:    queue.New[string](),
	}
	logger.Layout = logger.LogLayout
	logger.SetColor(cell.ColorWhite)
	return logger
}
func (l *Logger) SetColor(color cell.Color) {
	l.color = color
	l.ResetWidgets()
}
func (l *Logger) Add(msg string) {
	l.Data.Enqueue(msg)
	data := format.Text{
		Text: fmt.Sprintf("%s\n", msg),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	l.Display.Write(data.Text, data.Opts...)
	if l.Data.Len() >= maxCap {
		l.Data.Dequeue()
	}
}
func (l *Logger) AddF(msg string) {
	l.Data.Enqueue(msg)
	time := format.Text{
		Text: time.Now().Format(timeFormat),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	name := format.Text{
		Text: fmt.Sprintf(" [%s] ", strings.ToUpper(l.name)),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(l.color))},
	}
	data := format.Text{
		Text: fmt.Sprintf("%s\n", msg),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	l.Display.Write(time.Text, time.Opts...)
	l.Display.Write(name.Text, name.Opts...)
	l.Display.Write(data.Text, data.Opts...)
}
