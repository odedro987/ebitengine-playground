package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/pkg/sprite"
)

type player struct {
	sprite.Base
	speed float64
}

func NewPlayer(x, y float64) *player {
	p := &player{
		speed: 60,
	}
	p.SetPosition(x, y)
	p.LoadGraphic("assets/WispAvatar.png")

	return p
}

func (p *player) Update(elapsed float64) error {
	err := p.Base.Update(elapsed)
	if err != nil {
		return err
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X += p.speed * elapsed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X -= p.speed * elapsed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= p.speed * elapsed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += p.speed * elapsed
	}

	return nil
}
