package main

import (
	"flag"
	"time"

	"github.com/danmane/abalone/go/game"
	"github.com/danmane/abalone/go/quickstart"
)

var (
	port = flag.String("port", "3423", "port the ai runs on")
)

func play(s game.State, limit time.Duration) game.State {
	f := s.Futures()
	if len(f) == 0 {
		panic("Ah! There are no future states. Why'd the server send this to me? =( ")
	}
	return *f[0]
}

func main() {
	flag.Parse()
	quickstart.Run(*port, play)
}
