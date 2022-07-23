package main

import (
	_ "image/png"

	"github.com/odedro987/ebitengine-playground/pkg/game"
	"github.com/odedro987/ebitengine-playground/pkg/state"
)

func main() {
    s := state.MigStateBase{}

    game.NewGame(640, 480, "MIG", &s)
}