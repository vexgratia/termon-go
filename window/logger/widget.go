package logger

// This file contains the implementation of Logger widgets.

import (
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"
)

// reset recreates all Logger widgets.
//
// Blocks Logger mutex to avoid data race.
func (l *Logger) reset() {
	l.mu.Lock()
	l.settings = l.makeSettingsButton()
	l.log = l.makeLogButton()
	l.spark = l.makeSpark()
	l.display = l.makeDisplay()
	l.mu.Unlock()
}

// makeSpark creates Logger spark.
func (l *Logger) makeSpark() *sparkline.SparkLine {
	spark, _ := sparkline.New(
		sparkline.Label(" Tick: "),
		sparkline.Color(l.Color()),
	)
	return spark
}

// makeDisplay creates Logger display.
func (l *Logger) makeDisplay() *text.Text {
	text, _ := text.New(
		text.WrapAtWords(),
		text.RollContent(),
	)
	return text
}

// makeSettingsButton creates a button that switches Logger layout to settingsLayout.
func (l *Logger) makeSettingsButton() *button.Button {
	button, _ := button.New(
		"SET",
		func() error {
			l.setLayout(l.settingsLayout)
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(l.Color()),
	)
	return button
}

// makeLogButton creates a button that switches Logger layout to logLayout.
func (l *Logger) makeLogButton() *button.Button {
	button, _ := button.New(
		"LOG",
		func() error {
			l.setLayout(l.logLayout)
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(l.Color()),
	)
	return button
}
