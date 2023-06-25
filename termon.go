package termon

import (
	"context"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/vexgratia/termon-go/metric"
	storage "github.com/vexgratia/termon-go/storage"
)

const newTick = 100 * time.Millisecond

type TermonMode int

type TermonModeFunc func(t *Termon) []container.Option

const (
	TUI_DEFAULT TermonMode = iota
	TUI_SETTINGS
)

type Termon struct {
	Terminal *tcell.Terminal
	//
	Main    *container.Container
	Layouts map[TermonMode]TermonModeFunc
	Mode    TermonMode
	//
	Tick    time.Duration
	Storage *storage.Storage
}

func New(terminal *tcell.Terminal, tick time.Duration) *Termon {
	tui := &Termon{
		Terminal: terminal,
		Layouts: map[TermonMode]TermonModeFunc{
			TUI_DEFAULT: DefaultMode,
		},
		Tick: tick,
	}
	//
	tui.Main, _ = container.New(
		terminal, container.ID("MAIN"),
		container.Border(linestyle.Round),
		container.BorderTitle("TERMON"),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
	)
	//
	tui.Storage = storage.New(tick, metric.AllMetrics)
	return tui
}

func (t *Termon) Opts() []container.Option {
	return t.Layouts[t.Mode](t)
}
func DefaultMode(tui *Termon) []container.Option {
	return []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(),
					container.Right(),
				),
			),
			container.Bottom(
				container.SplitVertical(
					container.Left(),
					container.Right(),
				),
			),
		),
	}
}

func (t *Termon) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	quitFunc := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}
	go t.GetUpdates()
	termdash.Run(ctx, t.Terminal, t.Main, termdash.KeyboardSubscriber(quitFunc), termdash.RedrawInterval(t.Storage.Tick))
}

func (t *Termon) GetUpdates() {
	for {
		time.Sleep(t.Tick)
	}
}
