package window

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/vexgratia/termon-go/metric"
)

type Cell struct {
	Name    string
	Metric  *metric.Metric
	Color   cell.Color
	Display *text.Text
}

func (w *Window) MakeCell(metric *metric.Metric) *Cell {
	text, _ := text.New(
		text.WrapAtRunes(),
		text.DisableScrolling(),
	)
	return &Cell{
		Name:    metric.Name,
		Metric:  metric,
		Color:   w.Color,
		Display: text,
	}
}
func (w *Window) EmptyCell() *Cell {
	text, _ := text.New(
		text.WrapAtRunes(),
		text.DisableScrolling(),
	)
	return &Cell{
		Color:   w.Color,
		Display: text,
	}
}
func (w *Window) MakeCells() []*Cell {
	cells := []*Cell{}
	for _, metric := range w.Metrics {
		cell := w.MakeCell(metric)
		cells = append(cells, cell)
	}
	for len(cells) < 15 {
		cell := w.EmptyCell()
		cells = append(cells, cell)
	}
	return cells
}
func (c *Cell) Update() {
	c.Display.Reset()
	if c.Metric == nil {
		return
	}
	c.Display.Write(
		c.Metric.CurrentF(),
		text.WriteCellOpts(cell.FgColor(cell.ColorWhite)),
	)
}

func (c *Cell) Layout() []container.Option {
	if c.Metric == nil {
		return []container.Option{
			container.Border(linestyle.None),
			container.BorderTitle("EMPTY"),
			container.BorderTitleAlignCenter(),
			container.BorderColor(c.Color),
			container.FocusedColor(c.Color),
			container.PlaceWidget(c.Display),
		}
	}
	return []container.Option{
		container.ID(c.Name),
		container.Border(linestyle.Light),
		container.BorderTitle(" " + c.Metric.Tag() + " "),
		container.BorderTitleAlignCenter(),
		container.BorderColor(c.Color),
		container.FocusedColor(c.Color),
		container.PlaceWidget(c.Display),
	}
}
