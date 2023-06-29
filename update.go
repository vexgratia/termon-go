package termon

func (t *Termon) Relayout() {
	t.Main.Update("MAIN", t.Opts()...)
}
func (t *Termon) SetLayout(layout LayoutFunc) {
	t.Layout = layout
	t.Relayout()
}
func (t *Termon) SetWindow(name string, index int) {
	t.Windows[index] = t.WindowMap[name]
}
func (t *Termon) GetUpdates() {
	for _, w := range t.Windows {
		go w.GetUpdates()
	}
	for {
		select {
		case msg := <-t.Signal:
			t.Main.Update(msg.Name, msg.Opts...)
		default:
		}
	}
}
