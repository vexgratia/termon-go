package termon

import (
	"context"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/update"
	"github.com/vexgratia/termon-go/window"
)

var tick = 5 * time.Millisecond

type Termon struct {
	Terminal *tcell.Terminal
	Main     *container.Container
	Layout   LayoutFunc
	//
	CPU    *window.Window
	GC     *window.Window
	Golang *window.Window
	Memory *window.Window
	//
	Updates chan update.Message
}

func New(terminal *tcell.Terminal) *Termon {
	termon := &Termon{
		Terminal: terminal,
		Updates:  make(chan update.Message, 100),
	}
	//
	termon.Main = termon.MakeMain()
	//
	termon.CPU = window.New("CPU", cell.ColorRed, metric.CPU, termon.Updates)
	termon.GC = window.New("GC", cell.ColorGreen, metric.GC, termon.Updates)
	termon.Golang = window.New("Golang", cell.ColorBlue, metric.Golang, termon.Updates)
	termon.Memory = window.New("Memory", cell.ColorYellow, metric.Memory, termon.Updates)
	//
	termon.Layout = termon.DefaultLayout
	termon.Update()
	return termon
}

func (t *Termon) Opts() []container.Option {
	return t.Layout()
}

func (t *Termon) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyEsc {
			cancel()
		}
	}
	go t.GetUpdates()
	termdash.Run(ctx, t.Terminal, t.Main, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(tick))
}
