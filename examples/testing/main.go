package main

import (
	"github.com/odedro987/gixel-engine/examples/testing/states"
	"github.com/odedro987/gixel-engine/gixel"
)

func main() {
	s := states.PlayState{}

	gixel.NewGame(640, 480, "MIG", &s)
}
