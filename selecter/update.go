package selecter

func (s *Selecter) Update() {
	s.Updater.Request(s)
}
