# Conway's Game of Life in Go
=============

Implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life) in an infinite space in Go. Alive cells are stored in a set. To calculate the next iteration, we compute the number of neighbors for each cell that has neighbors.


## Purpose

I implemented this game to try Go and compare it with Python. My goal was to write as idiomatic as possible although I had no experience with Go whatsoever and I was on a place without Internet while writing this.


## Usage

To use the library, write a program like:

```go
package main

import (
	"fmt"
	"gol"
)

func main() {
	board := gol.MakeBoard(".X.\n..X\nXXX")
	for i := 0; i < 10; i++ {
		board.Advance()
		fmt.Println(board.DebugString())
	}
}
```