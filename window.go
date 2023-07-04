package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/vexgratia/termon-go/logger"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/tracker"
)

var gridCap = 16

func (t *Termon) Name() string {
	return "Main"
}
func (t *Termon) Color() cell.Color {
	return cell.ColorWhite
}

func (t *Termon) Opts() []container.Option {
	return t.Layout()
}
func (t *Termon) MakeWindows() {
	//
	cpu := tracker.New("CPU", t.Updater)
	cpu.Add(metric.GetColored(metric.CPU, cell.ColorRed)...)
	cpu.SetColor(cell.ColorRed)
	//
	gc := tracker.New("GC", t.Updater)
	gc.Add(metric.GetColored(metric.GC, cell.ColorGreen)...)
	gc.SetColor(cell.ColorGreen)
	//
	golang := tracker.New("Golang", t.Updater)
	golang.Add(metric.GetColored(metric.Golang, cell.ColorBlue)...)
	golang.SetColor(cell.ColorBlue)
	//
	memory := tracker.New("Memory", t.Updater)
	memory.Add(metric.GetColored(metric.Memory, cell.ColorYellow)...)
	memory.SetColor(cell.ColorYellow)
	//
	infoLog := logger.New("Info", t.Updater)
	infoLog.SetColor(cell.ColorBlue)
	t.Logger["Info"] = infoLog
	//
	warnLog := logger.New("Warn", t.Updater)
	warnLog.SetColor(cell.ColorYellow)
	t.Logger["Warn"] = warnLog
	//
	errLog := logger.New("Error", t.Updater)
	errLog.SetColor(cell.ColorRed)
	t.Logger["Error"] = errLog
	t.Add(golang, memory, cpu, gc, infoLog, warnLog, errLog)
}
