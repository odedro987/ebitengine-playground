package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/font"
)

const DISTANCE_FROM_WALL = 64

type PlayState struct {
	gixel.BaseGxlState
	paddle1, paddle2 *Paddle
	score1, score2   int
	scoreText        gixel.GxlText
	paddles          gixel.GxlGroup
	ball             *Ball
}

func (s *PlayState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	s.paddles = gixel.NewGroup(0)
	s.Add(s.paddles)

	s.paddle1 = NewPaddle(DISTANCE_FROM_WALL, GAME_HEIGHT/2)
	s.paddles.Add(s.paddle1)
	s.paddle1.SetControls(ebiten.KeyW, ebiten.KeyS)

	s.paddle2 = NewPaddle(GAME_WIDTH-DISTANCE_FROM_WALL-float64(*s.paddle1.W()), GAME_HEIGHT/2)
	s.paddles.Add(s.paddle2)
	s.paddle2.SetControls(ebiten.KeyUp, ebiten.KeyDown)

	s.ball = NewBall()
	s.Add(s.ball)

	f := font.NewFont("assets/nokiafc22.ttf")
	s.scoreText = gixel.NewText(0, 0, "0 - 0", f.GetPreset(64))

	s.Add(s.scoreText)

}

func (s *PlayState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	s.OverlapsObjectGroup(s.ball, s.paddles, func(obj1, obj2 *gixel.GxlObject) { s.ball.FlipHorizontal() })

	s.scoreText.SetText(fmt.Sprintf("%d - %d", s.score1, s.score2))
	s.scoreText.SetPosition(GAME_WIDTH/2-float64(*s.scoreText.W()/2), 64)

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		fullscreen := ebiten.IsFullscreen()
		fullscreen = !fullscreen
		ebiten.SetFullscreen(fullscreen)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		*s.Game().Timescale() = 0.5
	}

	return nil
}
