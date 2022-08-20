package main

import (
	"image/color"

	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/systems/physics"
)

type Paddle struct {
	gixel.BaseGxlSprite
	// Systems
	physics.Physics
	Movement
}

func NewPaddle(x, y float64) *Paddle {
	w := &Paddle{}
	w.SetPosition(x, y)

	return w
}

func (p *Paddle) Init(game *gixel.GxlGame) {
	p.BaseGxlSprite.Init(game)
	p.Physics.Init(p)
	p.Movement.Init(p)

	p.MakeGraphic(48, GAME_HEIGHT/6, color.White)
}

func (p *Paddle) Update(elapsed float64) error {
	err := p.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	p.Physics.Update(elapsed)
	p.Movement.Update(elapsed)

	return nil
}
