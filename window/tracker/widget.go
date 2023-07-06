package tracker

// This file contains the implementation of Tracker widgets.

import (
	"github.com/mum4k/termdash/widgets/button"
)

// reset recreates all Tracker widgets.
//
// Blocks Tracker mutex to avoid data race.
func (t *Tracker) reset() {
	t.mu.Lock()
	t.settings = t.makeSettingsButton()
	t.chart = t.makeFocusButton()
	t.cell = t.makeCellButton()
	t.mu.Unlock()
}

// makeSettingsButton creates a button that switches Tracker layout to settingsLayout.
func (t *Tracker) makeSettingsButton() *button.Button {
	button, _ := button.New(
		"SET",
		func() error {
			t.setLayout(t.settingsLayout)
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(t.Color()),
	)
	return button
}

// makeSettingsButton creates a button that switches Tracker layout to focusLayout.
func (t *Tracker) makeFocusButton() *button.Button {
	button, _ := button.New(
		"FOCUS",
		func() error {
			t.setLayout(t.focusLayout)
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(t.Color()),
	)
	return button
}

// makeCellButton creates a button that switches Tracker layout to cellLayout.
func (t *Tracker) makeCellButton() *button.Button {
	button, _ := button.New(
		"CELL",
		func() error {
			t.setLayout(t.cellLayout)
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(t.Color()),
	)
	return button
}
