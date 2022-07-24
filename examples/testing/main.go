package main

import (
	"github.com/odedro987/gixel-engine/examples/testing/states"
	"github.com/odedro987/gixel-engine/pkg/game"
)

func main() {
	s := states.PlayState{}

	game.NewGame(640, 480, "MIG", &s)
}
