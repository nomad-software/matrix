package main

import (
	"time"

	"github.com/nomad-software/matrix/cli"
	"github.com/nomad-software/matrix/matrix"
	"github.com/nomad-software/matrix/term"
)

func main() {
	opt := cli.ParseOptions()

	term := term.New(opt)
	view := matrix.New(term.Width, term.Height)

	quit := make(chan bool)
	term.HandleInput(quit)

loop:
	for {
		view.Iterate()
		term.Draw(view)
		term.Update()

		select {
		case <-quit:
			break loop
		case <-time.After(time.Second / 20):
		}
	}
	term.Destroy()
}
