package termon

func (t *Termon) Relayout() {
	t.Updater.SetWindow("MAIN", t.Layout())
}
func (t *Termon) SetLayout(layout LayoutFunc) {
	t.Layout = layout
	t.Relayout()
}
