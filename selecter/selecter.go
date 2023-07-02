package selecter

import (
	"github.com/mum4k/termdash/widgets/button"
	"github.com/vexgratia/termon-go/updater"
	"github.com/vexgratia/termon-go/window"
)

type Selecter struct {
	ID      int
	Windows []window.Window
	Buttons []*button.Button
	Updater *updater.Updater
}

func New(id int, updater *updater.Updater) *Selecter {
	return &Selecter{
		ID:      id,
		Updater: updater,
	}
}
func (s *Selecter) Add(w window.Window) {
	s.Windows = append(s.Windows, w)
	s.Buttons = s.MakeButtons()
}
func (s *Selecter) WindowSetFunc(w window.Window) func() error {
	return func() error {
		s.Updater.SetWindow(s.Name(), w.Opts())
		return nil
	}
}
