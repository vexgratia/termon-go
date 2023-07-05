package scroller

func (s *Scroller[T]) Update() {
	textChunk := s.Formatter(s.Current())
	s.Display.Reset()
	for _, text := range textChunk {
		s.Display.Write(text.Text, text.Opts...)
	}
}
