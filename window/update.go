package window

import "github.com/vexgratia/termon-go/update"

func (w *Window) Update() {
	w.Updates <- update.Message{
		Name: w.Name,
		Opts: w.Opts(),
	}
}
