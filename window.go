package termon

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
)

type WindowMode int

type WindowModeFunc func(w *Window) []container.Option

const (
	WINDOW_DEFAULT WindowMode = iota
	WINDOW_SETTINGS
)

type Window struct {
	TUI     *TUI
	Name    string
	Color   cell.Color
	Layouts map[WindowMode]WindowModeFunc
	Mode    WindowMode
	Metrics []string
	Tracked string
	Buttons []*button.Button
	Chart   *linechart.LineChart
}

func (tui *TUI) InitWindow(name string, color cell.Color, metrics []string) *Window {
	window := &Window{
		TUI:   tui,
		Name:  name,
		Color: color,
		Layouts: map[WindowMode]WindowModeFunc{
			WINDOW_DEFAULT: DefaultMode,
		},
		Metrics: metrics,
		Tracked: metrics[0],
	}
	window.Buttons = window.MakeButtons()
	window.Chart = window.MakeChart()
	window.SetMode()
	return window
}

func (w *Window) MakeButtons() []*button.Button {
	buttons := []*button.Button{}
	for _, name := range w.Metrics {
		b, _ := button.New(
			name,
			func() error {
				return nil
			},
			button.Height(1),
			button.DisableShadow(),
			button.FillColor(w.Color),
		)
		buttons = append(buttons, b)
	}
	return buttons
}

func (w *Window) MakeChart() *linechart.LineChart {
	chart, _ := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorWhite)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorWhite)),
	)
	return chart
}

func (w *Window) SetMode() {
	w.TUI.Main.Update(w.Name, w.Opts()...)
}
func (w *Window) Opts() []container.Option {
	return w.Layouts[w.Mode](w)
}

func DefaultMode(w *Window) []container.Option {
	return []container.Option{
		container.ID(w.Name),
		container.Border(linestyle.Round),
		container.BorderTitle(w.Name),
		container.BorderTitleAlignCenter(),

		container.BorderColor(w.Color),
		container.FocusedColor(w.Color),
		container.SplitHorizontal(
			container.Top(
				container.ID(w.Name+"_BUTTONS"),
			),
			container.Bottom(
				container.ID(w.Name+"_CHART"),
				container.PlaceWidget(w.Chart),
			),
			container.SplitPercent(30),
		),
	}
}

func (w *Window) UpdateChart() {
	w.Chart.Series("serie", w.TUI.Storage.Metrics[w.Tracked].Queue.Data, linechart.SeriesCellOpts(cell.FgColor(w.Color)))
}
func (w *Window) Update() {
	w.UpdateChart()
}
