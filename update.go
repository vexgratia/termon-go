package termon

import "github.com/mum4k/termdash/container"

type UpdateMessage struct {
	Name string
	Opts []container.Option
}

func (t *Termon) Update() {
	t.Updates <- UpdateMessage{
		Name: "MAIN",
		Opts: t.Opts(),
	}
}
func (t *Termon) GetUpdates() {
	for {
		select {
		case msg := <-t.Updates:
			t.Main.Update(msg.Name, msg.Opts...)
		default:
			continue
		}
	}
}
