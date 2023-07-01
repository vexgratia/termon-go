package termon

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/vexgratia/termon-go/updater"
	"github.com/vexgratia/termon-go/window"
)

var tick = 50 * time.Millisecond

type Termon struct {
	Terminal *tcell.Terminal
	mu       *sync.Mutex
	Main     *container.Container
	Layout   LayoutFunc
	Windows  map[string]window.Window
	Pool     map[string]window.Window
	Counter  atomic.Uint32
	Updater  *updater.Updater
}

func New(terminal *tcell.Terminal) *Termon {
	termon := &Termon{
		Terminal: terminal,
		mu:       &sync.Mutex{},
		Pool:     make(map[string]window.Window),
		Windows:  make(map[string]window.Window),
	}
	termon.Main = termon.MakeMain()
	termon.Updater = updater.New(termon.Main)
	return termon
}

func (t *Termon) Add(windows ...window.Window) {
	for _, w := range windows {
		t.Windows[w.Name()] = w
		// t.Pool[w.Name()] = w
		go w.Run()
	}
}

func (t *Termon) Run() {
	ctx := context.Background()
	t.MakeWindows()
	switcher := func(k *terminalapi.Keyboard) {
		switch k.Key {
		case '1':
			t.SetLayout(t.FocusLayout)
		case '2':
			t.SetLayout(t.SplitLayout)
		default:
		}
	}
	termdash.Run(ctx, t.Terminal, t.Main, termdash.KeyboardSubscriber(switcher), termdash.KeyboardSubscriber(switcher), termdash.RedrawInterval(tick))
}
