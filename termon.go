package termon

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/vexgratia/termon-go/selecter"
	"github.com/vexgratia/termon-go/updater"
	"github.com/vexgratia/termon-go/window"
)

var tick = 50 * time.Millisecond
var maxWindows = 8

type Termon struct {
	Terminal  *tcell.Terminal
	mu        *sync.Mutex
	Main      *container.Container
	Selecters []*selecter.Selecter
	Layout    LayoutFunc
	Updater   *updater.Updater
}

func New(terminal *tcell.Terminal) *Termon {
	termon := &Termon{
		Terminal: terminal,
		mu:       &sync.Mutex{},
	}
	main, _ := container.New(
		termon.Terminal,
		container.ID("Termon"),
		container.Border(linestyle.Round),
		container.BorderTitle(" TERMON "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
		container.SplitHorizontal(
			container.Top(),
			container.Bottom(container.ID("Main")),
			container.SplitPercent(1),
		),
	)
	termon.Main = main
	termon.Updater = updater.New(termon.Main)
	for id := 0; id < maxWindows; id++ {
		termon.Selecters = append(termon.Selecters, selecter.New(id, termon.Updater))
	}
	termon.MakeWindows()
	return termon
}

func (t *Termon) Add(windows ...window.Window) {
	for _, w := range windows {
		for _, s := range t.Selecters {
			s.Add(w)
			s.Update()
		}
		go w.Run()
	}
}
func (t *Termon) Run() {

}
func Run(ctx context.Context) {
	terminal, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer terminal.Close()
	t := New(terminal)
	ctx, cancel := context.WithCancel(ctx)
	switcher := func(k *terminalapi.Keyboard) {
		switch k.Key {
		case '1':
			t.SetLayout(t.FocusLayout)
		case '2':
			t.SetLayout(t.SplitLayout)
		case '3':
			t.SetLayout(t.TrioLayout)
		case '4':
			t.SetLayout(t.QuadroLayout)
		case 'q':
			cancel()
		default:
		}
	}
	fmt.Println("AWDAWDAWD")
	termdash.Run(ctx, t.Terminal, t.Main, termdash.KeyboardSubscriber(switcher), termdash.KeyboardSubscriber(switcher), termdash.RedrawInterval(tick))
}
