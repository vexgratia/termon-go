package termon

import "github.com/mum4k/termdash/container"

type TermonLayout int

type LayoutFunc func() []container.Option

const (
	TERMON_DEFAULT TermonLayout = iota
	TERMON_SETTINGS
)

func (t *Termon) DefaultLayout() []container.Option {
	return []container.Option{
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(t.CPU.Opts()...),
					container.Right(t.GC.Opts()...),
				),
			),
			container.Bottom(
				container.SplitVertical(
					container.Left(t.Golang.Opts()...),
					container.Right(t.Memory.Opts()...),
				),
			),
		),
	}
}
