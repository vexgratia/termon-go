package termon

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
)

type LayoutFunc func() []container.Option

func (t *Termon) FocusLayout() []container.Option {
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(),
			container.Bottom(t.Selecter[0].Opts()...),
			container.SplitPercent(1),
		),
	}
	return opts
}

func (t *Termon) SplitLayout() []container.Option {
	opts := []container.Option{
		container.SplitHorizontal(
			container.Top(t.Selecter[0].Opts()...),
			container.Bottom(t.Selecter[1].Opts()...),
			container.SplitPercent(50),
		),
	}
	return opts
}
func (t *Termon) TrioLayout() []container.Option {
	builder := grid.New()
	builder.Add(
		grid.ColWidthPercWithOpts(34, t.Selecter[0].Opts()),
		grid.ColWidthPercWithOpts(34, t.Selecter[1].Opts()),
		grid.ColWidthPercWithOpts(32, t.Selecter[2].Opts()),
	)
	opts, _ := builder.Build()
	return opts
}
func (t *Termon) QuadroLayout() []container.Option {
	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(50,
			grid.ColWidthPercWithOpts(50, t.Selecter[0].Opts()),
			grid.ColWidthPercWithOpts(50, t.Selecter[1].Opts()),
		),
	)
	builder.Add(
		grid.RowHeightPerc(50,
			grid.ColWidthPercWithOpts(50, t.Selecter[2].Opts()),
			grid.ColWidthPercWithOpts(50, t.Selecter[3].Opts()),
		),
	)
	opts, _ := builder.Build()
	return opts
}
