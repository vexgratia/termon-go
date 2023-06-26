package termon

import "github.com/vexgratia/termon-go/update"

func (t *Termon) Update() {
	t.Updates <- update.Message{
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
