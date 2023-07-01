package updater

import (
	"github.com/mum4k/termdash/container"
	"github.com/vexgratia/termon-go/window"
)

var buf = 100

type Updater struct {
	Main *container.Container
}
type Message struct {
	Name string
	Opts []container.Option
}

func New(main *container.Container) *Updater {
	return &Updater{
		Main: main,
	}
}
func (u *Updater) Request(w window.Window) {
	w.Lock()
	u.Main.Update(w.Name(), w.Opts()...)
	w.Unlock()
}

func (u *Updater) SetWindow(id string, opts []container.Option) {
	u.Main.Update(id, opts...)
}
