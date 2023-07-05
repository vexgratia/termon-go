package selecter

import (
	"strconv"

	"github.com/mum4k/termdash/cell"
)

func (s *Selecter) Name() string {
	return strconv.Itoa(s.ID)
}

func (s *Selecter) Color() cell.Color {
	return cell.ColorWhite
}

func (s *Selecter) Run() {

}
