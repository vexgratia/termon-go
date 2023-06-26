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
	"github.com/vexgratia/termon-go/cache"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/update"
	"github.com/vexgratia/termon-go/window"
)

type Termon struct {
	Terminal  *tcell.Terminal
	Main      *container.Container
	Layout    TermonLayout
	LayoutSet map[TermonLayout]LayoutFunc
	//
	CPU    *window.Window
	GC     *window.Window
	Golang *window.Window
	Memory *window.Window
	//
	Tick  time.Duration
	Cache *cache.Cache
	//
	Updates chan update.Message
}

func New(terminal *tcell.Terminal, tick time.Duration) *Termon {
	termon := &Termon{
		Terminal: terminal,
		Tick:     tick,
		Layout:   TERMON_DEFAULT,
		Updates:  make(chan update.Message, 100),
	}
	//
	termon.LayoutSet = map[TermonLayout]LayoutFunc{
		TERMON_DEFAULT: termon.DefaultLayout,
	}
	//
	termon.Cache = cache.New()
	//
	termon.Main = termon.MakeMain()
	//
	termon.CPU = window.New("CPU", cell.ColorRed, termon.Cache.GetMetrics(metric.CPU), termon.Updates)
	termon.GC = window.New("GC", cell.ColorGreen, termon.Cache.GetMetrics(metric.GC), termon.Updates)
	termon.Golang = window.New("Golang", cell.ColorBlue, termon.Cache.GetMetrics(metric.Golang), termon.Updates)
	termon.Memory = window.New("Memory", cell.ColorYellow, termon.Cache.GetMetrics(metric.Memory), termon.Updates)
	//
	return termon
}

func (t *Termon) Opts() []container.Option {
	return t.LayoutSet[t.Layout]()
}

func (t *Termon) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyEsc {
			cancel()
		}
	}
	go t.GetUpdates()
	termdash.Run(ctx, t.Terminal, t.Main, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(t.Tick))
}
