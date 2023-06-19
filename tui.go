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
)

const newTick = 10 * time.Millisecond

type TUIMode int

type TUIModeFunc func(t *TUI) []container.Option

const (
	TUI_DEFAULT TUIMode = iota
	TUI_SETTINGS
)

var SettingsButtons = map[string]rune{
	"CPU_SET":    '1',
	"GC_SET":     '2',
	"Other_SET":  '3',
	"Memory_SET": '4',
}

type TUI struct {
	Term *tcell.Terminal
	//
	Main    *container.Container
	Layouts map[TUIMode]TUIModeFunc
	Mode    TUIMode
	//
	CPU    *Window
	GC     *Window
	Other  *Window
	Memory *Window
	//
	Tick    time.Duration
	Storage *Storage
}

func InitTUI(term *tcell.Terminal, tick time.Duration) *TUI {
	tui := &TUI{
		Term: term,
		Layouts: map[TUIMode]TUIModeFunc{
			TUI_DEFAULT: DefaulMode,
		},
		Tick: tick,
	}
	//
	tui.Main, _ = container.New(term, container.ID("MAIN"),
		container.Border(linestyle.Round),
		container.BorderTitle("TERMON"),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite))
	//
	tui.InitStorage()
	tui.CPU = tui.InitWindow("CPU", cell.ColorRed, CPUMetrics)
	tui.GC = tui.InitWindow("GC", cell.ColorGreen, GCMetrics)
	tui.Other = tui.InitWindow("Other", cell.ColorBlue, OtherMetrics)
	tui.Memory = tui.InitWindow("Memory", cell.ColorYellow, MemoryMetrics)
	tui.SetMode()
	return tui
}
func (tui *TUI) SetMode() {
	tui.Main.Update("MAIN", tui.Opts()...)
}
func (tui *TUI) Opts() []container.Option {
	return tui.Layouts[tui.Mode](tui)
}
func DefaulMode(tui *TUI) []container.Option {
	return []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(
						tui.CPU.Opts()...,
					),
					container.Right(
						tui.GC.Opts()...,
					),
				),
			),
			container.Bottom(
				container.SplitVertical(
					container.Left(
						tui.Other.Opts()...,
					),
					container.Right(
						tui.Memory.Opts()...,
					),
				),
			),
		),
	}
}

func (tui *TUI) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	quitFunc := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}
	go tui.Storage.GetUpdates()
	go tui.GetUpdates()
	termdash.Run(ctx, tui.Term, tui.Main, termdash.KeyboardSubscriber(quitFunc), termdash.RedrawInterval(tui.Storage.Tick))
}
func (tui *TUI) GetUpdates() {
	for {
		tui.CPU.Update()
		tui.GC.Update()
		tui.Other.Update()
		tui.Memory.Update()
		time.Sleep(tui.Tick)
	}
}
