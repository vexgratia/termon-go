package format

import (
	"github.com/mum4k/termdash/widgets/text"
)

type TextWithOpts struct {
	Text string
	Opts []text.WriteOption
}
