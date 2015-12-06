# Conway's Game of Life in Go

Implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life) in an infinite space in Go. Alive cells are stored in a set. To calculate the next iteration, we compute the number of neighbors for each cell that has neighbors.


## Purpose

I implemented this game to try Go and compare it with Python. My goal was to write as idiomatic as possible although I had no experience with Go whatsoever and I was on a plane without Internet while writing this.

Check out [the example](https://github.com/domoritz/gameoflife-go/blob/master/golapp/app.go).

## Help

### Library

Change into the app directory with `cd gol`. Run tests with `go test -v`.

### App

Change into the app directory with `cd golapp`. Run the main program with `go run app.go`. Build it with `go build`.
