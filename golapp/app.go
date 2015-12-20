package main

import (
	"fmt"
	"time"

	"github.com/domoritz/gameoflife-go/gol"
)

func main() {
	// Diehard
	field := gol.MakeField("......X.\nXX......\n.X...XXX")
	for i := 0; i < 130; i++ {
		field = field.Step()
		fmt.Print("\033[2J\033[1;1H", field)
		time.Sleep(time.Second / 10)
	}
}
