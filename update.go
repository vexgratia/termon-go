package termon

func (t *Termon) Relayout() {
	t.Updater.Request(t)
}
func (t *Termon) SetLayout(layout LayoutFunc) {
	t.Layout = layout
	t.Relayout()
}
