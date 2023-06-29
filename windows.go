package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/vexgratia/termon-go/metric"
	"github.com/vexgratia/termon-go/tracker"
)

type Window interface {
	Name() string
	Opts() []container.Option
	GetUpdates()
}

func (t *Termon) MakeWindows() {
	//
	cpu := tracker.New("CPU")
	cpu.Add(metric.GetColored(metric.CPU, cell.ColorRed)...)
	cpu.SetColor(cell.ColorRed)
	cpu.Connect(t.Signal)
	//
	gc := tracker.New("GC")
	gc.Add(metric.GetColored(metric.GC, cell.ColorGreen)...)
	gc.SetColor(cell.ColorGreen)
	gc.Connect(t.Signal)
	//
	golang := tracker.New("Golang")
	golang.Add(metric.GetColored(metric.Golang, cell.ColorBlue)...)
	golang.SetColor(cell.ColorBlue)
	golang.Connect(t.Signal)
	//
	memory := tracker.New("Memory")
	memory.Add(metric.GetColored(metric.Memory, cell.ColorYellow)...)
	memory.SetColor(cell.ColorYellow)
	memory.Connect(t.Signal)
	t.Add(cpu, gc, golang, memory)
}
