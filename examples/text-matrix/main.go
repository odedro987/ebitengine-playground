package main

import (
	"github.com/odedro987/gixel-engine/gixel"
)

const GAME_WIDTH = 1280
const GAME_HEIGHT = 960

func main() {
	gixel.NewGame(1280, 960, "Matrix", &PlayState{}, 1)
}
