package main

import (
	"github.com/odedro987/gixel-engine/examples/testing/states"
	"github.com/odedro987/gixel-engine/gixel"
)

func main() {
	gixel.NewGame(640, 480, "Hello Gixel", &states.MenuState{}, 2)
}
