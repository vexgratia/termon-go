package termon

import (
	"strconv"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/vexgratia/termon-go/scroller"
	"github.com/vexgratia/termon-go/window"
)

type LayoutFunc func() []container.Option

func (t *Termon) FocusLayout() []container.Option {
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(),
			container.Bottom(t.ChooseWindowLayout(t.Counter.Add(1))...),
			container.SplitPercent(1),
		),
	}
	return opts
}
func (t *Termon) SplitLayout() []container.Option {
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(t.ChooseWindowLayout(t.Counter.Add(1))...),
			container.Bottom(t.ChooseWindowLayout(t.Counter.Add(1))...),
			container.SplitPercent(50),
		),
	}
	return opts
}

func (t *Termon) ChooseWindowLayout(id uint32) []container.Option {
	scr := scroller.New[window.Window]()
	for _, w := range t.Pool { // !!!
		scr.List.Push(w)
	}
	scr.Update() // !
	but, _ := button.New(
		"SUBMIT",
		func() error {
			conv := strconv.Itoa(int(id))
			t.Updater.SetWindow(conv, scr.Current().Opts())
			delete(t.Pool, scr.Current().Name())
			return nil
		},
		button.Height(2),
		button.Width(8),
		button.FillColor(t.Color()),
	)
	opts := []container.Option{
		container.ID(strconv.Itoa(int(id))),
		container.Border(linestyle.Round),
		container.BorderTitle(" CHOOSE WINDOW "),
		container.BorderTitleAlignCenter(),

		container.BorderColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorWhite),
		container.SplitHorizontal(
			container.Top(scr.Opts()...),
			container.Bottom(container.PlaceWidget(but)),
			container.SplitPercent(50),
		),
	}
	return opts
}
