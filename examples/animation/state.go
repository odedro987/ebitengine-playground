package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/odedro987/gixel-engine/gixel"
)

type PlayState struct {
	gixel.BaseGxlState
	player1 *Player
	player2 *Player
}

func (s *PlayState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	s.player1 = NewPlayer(50, 50)
	s.Add(s.player1)

	s.player2 = NewPlayer(50, 100)
	s.Add(s.player2)

	s.player1.SetControls(ebiten.KeyUp, ebiten.KeyDown, ebiten.KeyLeft, ebiten.KeyRight)
	s.player2.SetControls(ebiten.KeyW, ebiten.KeyS, ebiten.KeyA, ebiten.KeyD)

}

func (s *PlayState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		fullscreen := ebiten.IsFullscreen()
		fullscreen = !fullscreen
		ebiten.SetFullscreen(fullscreen)
	}

	return nil
}
