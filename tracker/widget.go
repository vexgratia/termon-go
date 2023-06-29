package tracker

import "github.com/mum4k/termdash/widgets/button"

func (t *Tracker) ResetWidgets() {
	t.Settings = t.MakeSettingsButton()
	t.Chart = t.MakeChartButton()
	t.Cell = t.MakeCellButton()
}

func (t *Tracker) MakeSettingsButton() *button.Button {
	button, _ := button.New(
		"SET",
		func() error {
			t.SetLayout(t.SettingsLayout)
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(t.Color),
	)
	return button
}
func (t *Tracker) MakeChartButton() *button.Button {
	button, _ := button.New(
		"CHART",
		func() error {
			t.SetLayout(t.ChartLayout)
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(t.Color),
	)
	return button
}
func (t *Tracker) MakeCellButton() *button.Button {
	button, _ := button.New(
		"CELL",
		func() error {
			t.SetLayout(t.CellLayout)
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(t.Color),
	)
	return button
}
