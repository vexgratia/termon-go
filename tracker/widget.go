package tracker

import "github.com/mum4k/termdash/widgets/button"

func (t *Tracker) ResetWidgets() {
	t.Settings = t.MakeSettingsButton()
}

func (t *Tracker) MakeSettingsButton() *button.Button {
	button, _ := button.New(
		"SET",
		func() error {
			t.Layout = t.ChartLayout
			t.Relayout()
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(t.Color),
	)
	return button
}
