package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/math"
)

type imports interface {
	Velocity() *math.GxlPoint
}

type Movement struct {
	upKey, downKey ebiten.Key
	speed          float64
	subject        *imports
}

func (p *Movement) Init(subject imports) {
	p.subject = &subject
	p.speed = 500
}

func (p *Movement) Update(elapsed float64) {
	p.keyboardControl()
}

func (p *Movement) SetControls(upKey, downKey ebiten.Key) {
	p.upKey, p.downKey = upKey, downKey
}

func (p *Movement) keyboardControl() {
	up := ebiten.IsKeyPressed(p.upKey)
	down := ebiten.IsKeyPressed(p.downKey)

	vel := 0.0

	if up {
		vel += -p.speed
	}

	if down {
		vel += p.speed
	}

	(*p.subject).Velocity().Y = vel
}
