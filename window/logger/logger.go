package logger

// This file contains the implementation of Logger and its basic methods.
import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/collection-go/generic/queue"
	"github.com/vexgratia/termon-go/format"
	"github.com/vexgratia/termon-go/template/scroller"
	"github.com/vexgratia/termon-go/updater"
)

var maxCap = 200

// A Tracker is a Window type that displays logs.
type Logger struct {
	// general
	name   string
	data   *queue.Queue[string]
	layout LayoutFunc
	// sync
	mu *sync.Mutex
	// templates
	color *scroller.Scroller[cell.Color]
	// widgets
	settings *button.Button
	log      *button.Button
	display  *text.Text
	spark    *sparkline.SparkLine
	// external
	updater *updater.Updater
}

// New creates a Logger based on Updater.
func New(name string, updater *updater.Updater) *Logger {
	l := &Logger{
		// general, sync, data and external
		name:    name,
		mu:      &sync.Mutex{},
		data:    queue.New[string](),
		updater: updater,
	}
	l.layout = l.logLayout
	// templates
	l.color = l.makeColorScroller()
	// widgets
	l.reset()
	return l
}

// Add appends message to Logger.
func (l *Logger) Add(msg string) {
	l.data.Enqueue(msg)
	data := format.Text{
		Text: fmt.Sprintf("%s\n", msg),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	l.display.Write(data.Text, data.Opts...)
	if l.data.Len() >= maxCap {
		l.data.Dequeue()
	}
}

// AddF formats and appends message to Logger.
func (l *Logger) AddF(msg string) {
	l.data.Enqueue(msg)
	time := format.Text{
		Text: time.Now().Format(timeFormat),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	name := format.Text{
		Text: fmt.Sprintf(" [%s] ", strings.ToUpper(l.name)),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(l.Color()))},
	}
	data := format.Text{
		Text: fmt.Sprintf("%s\n", msg),
		Opts: []text.WriteOption{text.WriteCellOpts(cell.FgColor(cell.ColorWhite))},
	}
	l.display.Write(time.Text, time.Opts...)
	l.display.Write(name.Text, name.Opts...)
	l.display.Write(data.Text, data.Opts...)
}

// Name returns Logger name.
func (l *Logger) Name() string {
	return l.name
}

// Color returns current Logger color.
func (l *Logger) Color() cell.Color {
	return l.color.Current()
}

// Opts configurates and returns opts based on current LayoutFunc.
func (l *Logger) Opts() []container.Option {
	opts := l.initOpts()
	opts = append(opts, l.layout()...)
	return opts
}
