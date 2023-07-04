package termon

func (t *Termon) Info(msg string) {
	t.Logger["Info"].AddF(msg)
}
func (t *Termon) Warn(msg string) {
	t.Logger["Warn"].AddF(msg)
}
func (t *Termon) Error(msg string) {
	t.Logger["Error"].AddF(msg)
}
