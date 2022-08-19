package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/systems/physics"
)

const BALL_SPEED = 500

type Ball struct {
	gixel.BaseGxlSprite

	// Systems
	physics.Physics
}

func NewBall() *Ball {
	b := &Ball{}

	return b
}

func (b *Ball) Init(game *gixel.GxlGame) {
	b.BaseGxlSprite.Init(game)
	b.Physics.Init(b)

	b.MakeGraphic(32, 32, color.White)

	b.Spawn()
}

func (b *Ball) Spawn() {
	b.SetPosition(GAME_WIDTH/2-16, GAME_HEIGHT/2-16)

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.Velocity().X = float64(BALL_SPEED)

	if randGen.Float64() > 0.5 {
		b.Velocity().X *= -1
	}

	b.Velocity().Y = float64(BALL_SPEED) * (randGen.Float64() - 0.5)
}

func (b *Ball) FlipVertical() {
	b.Velocity().Y *= -1
}

func (b *Ball) FlipHorizontal() {
	b.Velocity().X *= -1
}

func (b *Ball) Update(elapsed float64) error {
	err := b.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	b.Physics.Update(elapsed)

	if *b.Y() > GAME_HEIGHT-float64(*b.H()) || *b.Y() < 0 {
		b.FlipVertical()
	}

	state, _ := (*b.Game().State()).(*PlayState)

	if *b.X() > GAME_WIDTH-float64(*b.W()) {
		b.Spawn()
		state.score2++
	} else if *b.X() < 0 {
		b.Spawn()
		state.score1++
	}

	return nil
}
