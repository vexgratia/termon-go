package selecter

import "github.com/vexgratia/termon-go/window"

//
func (s *Selecter) Update() {

}

//
func (s *Selecter) GetUpdates() {

}

// relayout updates Selecter container with current opts.
func (s *Selecter) relayout() {
	s.updater.Request(s)
}

// windowSetFunc returns function that swaps Selecter with given Window.
func (s *Selecter) windowSetFunc(w window.Window) func() error {
	return func() error {
		s.updater.SetWindow(s.Name(), w.Opts())
		return nil
	}
}
