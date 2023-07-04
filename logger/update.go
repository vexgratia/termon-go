package logger

func (l *Logger) Update() {
	l.UpdateSpark()
}

func (l *Logger) UpdateSpark() {
	l.Spark.Add([]int{l.Data.Len()})
	l.Data.Clear()
}
