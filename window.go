package termon

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
	dcll "github.com/vexgratia/collection-go/models/generic/lists/circular/doubly"
)

type WindowMode int

type WindowLayout func(w *Window) []container.Option

const (
	WINDOW_DEFAULT WindowMode = iota
	WINDOW_SETTINGS
)

type Window struct {
	TUI      *TUI
	Name     string
	Color    cell.Color
	Layouts  map[WindowMode]WindowLayout
	Mode     WindowMode
	Metrics  []string
	Scroller *dcll.List[string]
	Row      *WindowRow
	Display  *WindowDisplay
}

func (w *Window) Update() {
	w.UpdateRow()
	w.UpdateDisplay()
}
func (w *Window) CurrentMetricName() string {
	return w.Scroller.Head.Value
}
func (w *Window) CurrentMetric() *Metric {
	return w.TUI.Storage.Metrics[w.CurrentMetricName()]
}
func (w *Window) Opts() []container.Option {
	return w.Layouts[w.Mode](w)
}

func (tui *TUI) InitWindow(name string, color cell.Color, metrics []string) *Window {
	window := &Window{
		TUI:   tui,
		Name:  name,
		Color: color,
		Layouts: map[WindowMode]WindowLayout{
			WINDOW_DEFAULT: DefaultLayout,
		},
		Metrics: metrics,
	}
	window.Scroller = window.MakeScroller()
	window.Row = window.InitRow()
	window.Display = window.InitDisplay()
	return window
}
func (w *Window) MakeScroller() *dcll.List[string] {
	scroller := dcll.New[string]()
	for _, metric := range w.Metrics {
		scroller.Push(metric)
	}
	return scroller
}

func DefaultLayout(w *Window) []container.Option {
	return []container.Option{
		container.ID(w.Name),
		container.Border(linestyle.Round),
		container.BorderTitle(w.Name),
		container.BorderTitleAlignCenter(),

		container.BorderColor(w.Color),
		container.FocusedColor(cell.ColorWhite),
		container.SplitHorizontal(
			container.Top(
				w.RowLayout()...,
			),
			container.Bottom(
				w.ChartLayout()...,
			),
			container.SplitPercent(30),
		),
	}
}

type WindowRow struct {
	Settings *button.Button
	Next     *button.Button
	Current  *text.Text
	Prev     *button.Button
}

func (w *Window) UpdateRow() {
	metric := w.CurrentMetric()
	name := fmt.Sprintf("%s: ", w.CurrentMetricName())
	data := fmt.Sprintf("%v\n", metric.Queue.Data[metric.Queue.Len()-1])
	w.Row.Current.Reset()
	w.Row.Current.Write(name, text.WriteCellOpts(cell.FgColor(cell.ColorWhite)))
	w.Row.Current.Write(data, text.WriteCellOpts(cell.FgColor(w.Color)))
}

func (w *Window) InitRow() *WindowRow {
	row := &WindowRow{
		Settings: w.MakeSettingsButton(),
		Current:  w.MakeCurrentText(),
	}
	row.Prev, row.Next = w.MakeScrollButtons()
	return row
}
func (w *Window) MakeSettingsButton() *button.Button {
	button, _ := button.New(
		"SET",
		func() error {
			w.Mode = WINDOW_SETTINGS
			return nil
		},
		button.Height(2),
		button.FillColor(w.Color),
	)
	return button
}

func (w *Window) MakeScrollButtons() (*button.Button, *button.Button) {
	prev, _ := button.New(
		"<---",
		func() error {
			w.Scroller.ScrollPrev()
			return nil
		},
		button.Height(2),
		button.FillColor(w.Color),
	)
	next, _ := button.New(
		"--->",
		func() error {
			w.Scroller.ScrollNext()
			return nil
		},
		button.Height(2),
		button.FillColor(w.Color),
	)
	return prev, next
}
func (w *Window) MakeCurrentText() *text.Text {
	current, _ := text.New(
		text.WrapAtRunes(),
	)
	return current
}
func (w *Window) RowLayout() []container.Option {
	builder := grid.New()
	builder.Add(
		grid.ColWidthPerc(15,
			grid.Widget(w.Row.Settings,
				container.Border(linestyle.None),
			),
		),
	)
	builder.Add(
		grid.ColWidthPerc(15,
			grid.Widget(w.Row.Prev,
				container.Border(linestyle.None),
			),
		),
	)
	builder.Add(
		grid.ColWidthPerc(55,
			grid.Widget(w.Row.Current,
				container.Border(linestyle.Round),
			),
		),
	)
	builder.Add(
		grid.ColWidthPerc(15,
			grid.Widget(w.Row.Next,
				container.Border(linestyle.None),
			),
		),
	)
	opts, _ := builder.Build()
	return opts
}

type WindowDisplay struct {
	Scroller *dcll.List[string]
	Chart    *linechart.LineChart
}

func (w *Window) UpdateDisplay() {
	w.Display.Chart.Series(
		"serie",
		w.CurrentMetric().Queue.Data,
		linechart.SeriesCellOpts(cell.FgColor(w.Color)),
	)
}
func (w *Window) InitDisplay() *WindowDisplay {
	display := &WindowDisplay{
		Chart: w.MakeChart(),
	}
	return display
}

func (w *Window) MakeChart() *linechart.LineChart {
	chart, _ := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YAxisAdaptive(),
	)
	return chart
}
func (w *Window) ChartLayout() []container.Option {
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(90,
			grid.Widget(w.Display.Chart,
				container.Border(linestyle.Round),
			),
		),
	)
	opts, _ := builder.Build()
	return opts
}

func (w *Window) MakeMetricButton(metric string) *button.Button {
	button, _ := button.New(
		metric,
		func() error {
			return nil
		},
		button.Height(2),
		button.FillColor(w.Color),
	)
	return button
}
