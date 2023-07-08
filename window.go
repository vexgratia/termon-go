package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/window/logger"
	"github.com/vexgratia/termon-go/window/tracker"
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
	cpu.Add(metric.Get(metric.CPU)...)
	//
	gc := tracker.New("GC", t.Updater)
	gc.Add(metric.Get(metric.GC)...)
	//
	golang := tracker.New("Golang", t.Updater)
	golang.Add(metric.Get(metric.Golang)...)
	//
	memory := tracker.New("Memory", t.Updater)
	memory.Add(metric.Get(metric.Memory)...)
	//
	infoLog := logger.New("Info", t.Updater)
	t.Logger["Info"] = infoLog
	//
	warnLog := logger.New("Warn", t.Updater)
	t.Logger["Warn"] = warnLog
	//
	errLog := logger.New("Error", t.Updater)
	t.Logger["Error"] = errLog
	//
	t.Add(golang, memory, cpu, gc, infoLog, warnLog, errLog)
}
