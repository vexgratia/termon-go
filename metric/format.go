package metric

func (m *Metric) ShortTag() string {
	var tag string
	for i := len(m.Name) - 1; i >= 0; i-- {
		char := m.Name[i]
		if char == '/' {
			break
		}
		tag = string(char) + tag
	}
	return tag
}
