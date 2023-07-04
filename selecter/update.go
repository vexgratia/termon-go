package selecter

func (s *Selecter) Update() {
	if !s.mu.TryLock() {
		return
	}
	s.Buttons = s.MakeButtons()
	s.mu.Unlock()
	// s.Updater.Request(s)
}
