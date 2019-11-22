# Conway's Game of Life in Go [![Build Status](https://travis-ci.org/domoritz/gameoflife-go.svg)](https://travis-ci.org/domoritz/gameoflife-go)

Implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life) in an infinite space in Go. Alive cells are stored in a set. To calculate the next iteration, we compute the number of neighbors for each cell that has neighbors.

I am implementing the Game of Life in different programming languages to learn about them. You can find [all of my implementations on GitHub](https://github.com/domoritz?tab=repositories&q=gameoflife).


## Purpose

I originally implemented this game to try Go and compare it with Python. My goal was to write as idiomatic as possible although I had no experience with Go whatsoever and I was on a plane without Internet while writing this.

Check out [the example](https://github.com/domoritz/gameoflife-go/blob/master/golapp/app.go).


### What I learned/liked/disliked about golang

* The translation form python was pretty straight forward. Nothing surprising.
* Not having list comprehensions or functional map makes finding the min and max a bit clunky but obvious.
* It's great to be able to extend a built in type.
* Why is there no proper set collection? `map[Cell]bool` is just awful.
* `go fmt` is just awesome. I don't think about formatting any more.
* Writing tests is super easy although I wish there was a simpler way to check for equality and automatically compute a diff.


## Help


### Library

Change into the app directory with `cd gol`. Run tests with `go test -v`.


### App

Change into the app directory with `cd golapp`. Run the main program with `go run app.go`. Build it with `go build`.
