package tracker

func (t *Tracker) Relayout() {
	t.Updater.Request(t)
}

func (t *Tracker) SetLayout(layout LayoutFunc) {
	t.Layout = layout
	t.Relayout()
}
