package termon

import (
	"context"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	cache "github.com/vexgratia/termon-go/cache"
	"github.com/vexgratia/termon-go/metric"
	window "github.com/vexgratia/termon-go/window"
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
	Tick    time.Duration
	Storage *cache.Cache
	//
	Updates chan UpdateMessage
}

func New(terminal *tcell.Terminal, tick time.Duration) *Termon {
	termon := &Termon{
		Terminal: terminal,
		Tick:     tick,
		Layout:   TERMON_DEFAULT,
	}
	//
	termon.LayoutSet = map[TermonLayout]LayoutFunc{
		TERMON_DEFAULT: termon.DefaultLayout,
	}
	//
	termon.Storage = cache.New(tick)
	//
	termon.Main, _ = container.New(
		terminal, container.ID("MAIN"),
		container.Border(linestyle.Round),
		container.BorderTitle("TERMON"),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
	)
	//
	termon.CPU = window.New("CPU", cell.ColorRed, termon.Storage.GetMetrics(metric.CPU))
	termon.GC = window.New("GC", cell.ColorGreen, termon.Storage.GetMetrics(metric.GC))
	termon.Golang = window.New("Golang", cell.ColorBlue, termon.Storage.GetMetrics(metric.Golang))
	termon.Memory = window.New("Memory", cell.ColorYellow, termon.Storage.GetMetrics(metric.Memory))
	//
	termon.Main.Update("MAIN", termon.Opts()...)
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
