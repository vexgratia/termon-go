package logger

import (
	"fmt"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/format"
	"github.com/vexgratia/termon-go/updater"
)

type Logger struct {
	InfoHandler *Window
	WarnHandler *Window
	ErrorHadler *Window
	Updater     *updater.Updater
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string) {
	time := format.TextWithOpts{
		Text: time.Now().Format(timeFormat),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	name := format.TextWithOpts{
		Text: " [INFO] ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(l.InfoHandler.Color))},
	}
	data := format.TextWithOpts{
		Text: fmt.Sprintf("%s\n", msg),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	l.InfoHandler.Display.Write(time.Text, time.Opts...)
	l.InfoHandler.Display.Write(name.Text, name.Opts...)
	l.InfoHandler.Display.Write(data.Text, data.Opts...)
}
func (l *Logger) Warn(msg string) {
	time := format.TextWithOpts{
		Text: time.Now().Format(timeFormat),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	name := format.TextWithOpts{
		Text: " [WARN] ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(l.WarnHandler.Color))},
	}
	data := format.TextWithOpts{
		Text: fmt.Sprintf("%s\n", msg),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	l.WarnHandler.Display.Write(time.Text, time.Opts...)
	l.WarnHandler.Display.Write(name.Text, name.Opts...)
	l.WarnHandler.Display.Write(data.Text, data.Opts...)
}
func (l *Logger) Error(msg string) {
	time := format.TextWithOpts{
		Text: time.Now().Format(timeFormat),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	name := format.TextWithOpts{
		Text: " [ERROR] ",
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(l.ErrorHadler.Color))},
	}
	data := format.TextWithOpts{
		Text: fmt.Sprintf("%s\n", msg),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	l.ErrorHadler.Display.Write(time.Text, time.Opts...)
	l.ErrorHadler.Display.Write(name.Text, name.Opts...)
	l.ErrorHadler.Display.Write(data.Text, data.Opts...)
}
