package main

import (
	"time"

	"github.com/mum4k/termdash/terminal/tcell"
	termon "github.com/vexgratia/termon-go"
)

var tick = time.Millisecond * 10

func main() {
	terminal, err := tcell.New()
	if err != nil {
		panic(err)
	}
	termon := termon.New(terminal, tick)
	termon.Run()
}
