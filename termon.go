package termon

import (
	"context"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/vexgratia/termon-go/update"
)

var tick = 5 * time.Millisecond

type Termon struct {
	Terminal  *tcell.Terminal
	Main      *container.Container
	Layout    LayoutFunc
	Windows   []Window
	WindowMap map[string]Window
	Signal    chan update.Message
}

func New(terminal *tcell.Terminal) *Termon {
	termon := &Termon{
		Terminal:  terminal,
		WindowMap: make(map[string]Window),
		Signal:    make(chan update.Message, 100),
	}
	termon.Main = termon.MakeMain()
	return termon
}
func (t *Termon) Name() string {
	return "MAIN"
}

func (t *Termon) Add(windows ...Window) {
	for _, w := range windows {
		t.Windows = append(t.Windows, w)
		t.WindowMap[w.Name()] = w
	}
}

func (t *Termon) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyEsc {
			cancel()
		}
	}
	t.MakeWindows()
	t.SetLayout(t.LayoutFour)
	go t.GetUpdates()
	termdash.Run(ctx, t.Terminal, t.Main, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(tick))
}
