package termon

func (t *Termon) Relayout() {
	t.Updater.Request(t)
}
func (t *Termon) SetLayout(layout LayoutFunc) {
	// for _, s := range t.Selecter {
	// 	s.Update()
	// }
	t.Layout = layout
	t.Relayout()
}
