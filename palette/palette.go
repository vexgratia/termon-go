package palette

import "github.com/mum4k/termdash/cell"

var (
	Maroon  = cell.ColorMaroon
	Green   = cell.ColorGreen
	Olive   = cell.ColorOlive
	Navy    = cell.ColorNavy
	Purple  = cell.ColorPurple
	Teal    = cell.ColorTeal
	Silver  = cell.ColorSilver
	Gray    = cell.ColorGray
	Red     = cell.ColorRed
	Lime    = cell.ColorLime
	Yellow  = cell.ColorYellow
	Blue    = cell.ColorBlue
	Fuchsia = cell.ColorFuchsia
	Aqua    = cell.ColorAqua
	White   = cell.ColorWhite
)
var All = []cell.Color{
	Maroon,
	Green,
	Olive,
	Navy,
	Purple,
	Teal,
	Silver,
	Gray,
	Red,
	Lime,
	Yellow,
	Blue,
	Fuchsia,
	Aqua,
	White,
}
var Palette = map[cell.Color]string{
	cell.ColorMaroon:  "Maroon",
	cell.ColorGreen:   "Green",
	cell.ColorOlive:   "Olive",
	cell.ColorNavy:    "Navy",
	cell.ColorPurple:  "Purple",
	cell.ColorTeal:    "Teal",
	cell.ColorSilver:  "Silver",
	cell.ColorGray:    "Gray",
	cell.ColorRed:     "Red",
	cell.ColorLime:    "Lime",
	cell.ColorYellow:  "Yellow",
	cell.ColorBlue:    "Blue",
	cell.ColorFuchsia: "Fuchsia",
	cell.ColorAqua:    "Aqua",
	cell.ColorWhite:   "White",
}

func String(c cell.Color) string {
	name, ok := Palette[c]
	if ok {
		return name
	}
	return ""
}
