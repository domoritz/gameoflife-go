package main

import (
	"fmt"
	"time"

	"github.com/domoritz/gameoflife-go/gol"
)

func main() {
	// Diehard
	f := gol.MakeField("......X.\nXX......\n.X...XXX")
	for i := 0; i < 130; i++ {
		f.Step()
		fmt.Print("\033[2J\033[1;1H", f.DebugString(2))
		time.Sleep(time.Second / 10)
	}
}
