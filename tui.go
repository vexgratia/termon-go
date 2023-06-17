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

const newTick = 100 * time.Millisecond

type TUI struct {
	Opts    *[]container.Option
	CPU     *CPU
	GC      *GC
	Other   *Other
	Memory  *Memory
	Storage *Storage
}

func NewTUI() *TUI {
	tui := &TUI{
		CPU:     NewCPU(),
		GC:      NewGC(),
		Other:   NewOther(),
		Memory:  NewMemory(),
		Storage: NewStorage(newTick),
	}
	tui.Opts = &[]container.Option{container.Border(linestyle.Round),
		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
		container.BorderTitle("TERMON"),
		container.ID("main"),
		container.BorderTitleAlignCenter(),
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(*tui.CPU.Opts...),
					container.Right(*tui.GC.Opts...),
				)),
			container.Bottom(
				container.SplitVertical(
					container.Left(*tui.Other.Opts...),
					container.Right(*tui.Memory.Opts...),
				)),
		),
	}
	return tui
}

func (tui *TUI) Run() {
	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()
	go tui.Storage.GetUpdates()
	main, err := container.New(
		t, *tui.Opts...,
	)
	ctx, cancel := context.WithCancel(context.Background())
	quit := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}
	termdash.Run(ctx, t, main, termdash.KeyboardSubscriber(quit), termdash.RedrawInterval(newTick))
}
