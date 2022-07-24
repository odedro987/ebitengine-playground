package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel"
)

type player struct {
	gixel.BaseGxlSprite
	speed float64
}

func NewPlayer(x, y float64) *player {
	p := &player{
		speed: 100,
	}
	p.SetPosition(x, y)

	return p
}

func (p *player) Init() {
	p.BaseGxlSprite.Init()
	p.LoadAnimatedGraphic("assets/Don2.png", 32, 28)
	p.SetFacingFlip(gixel.Right, false, false)
	p.SetFacingFlip(gixel.Left, true, false)
	p.Animation.Add("idle", []int{0}, 1, false)
	p.Animation.Add("walk", []int{2, 3, 4}, 15, true)
	p.Animation.Play("idle", false)
	// p.SetFacingFlip(gixel.Up, false, true)
	// p.SetFacingFlip(gixel.Down, false, false)
	// p.SetFacingFlip(gixel.Right|gixel.Up, true, true)
	// p.SetFacingFlip(gixel.Right|gixel.Down, true, false)
	// p.SetFacingFlip(gixel.Left|gixel.Up, false, true)
	// p.SetFacingFlip(gixel.Left|gixel.Down, false, false)
}

func (p *player) Update(elapsed float64) error {
	err := p.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	// facing := p.GetFacing()
	hasMoved := false

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X += p.speed * elapsed
		p.SetFacing(gixel.Right)
		hasMoved = true
		// facing |= gixel.Right
		// facing &= 0x1111 - gixel.Left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X -= p.speed * elapsed
		p.SetFacing(gixel.Left)
		hasMoved = true
		// facing |= gixel.Left
		// facing &= 0x1111 - gixel.Right
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= p.speed * elapsed
		hasMoved = true
		// facing |= gixel.Up
		// facing &= 0x1111 - gixel.Down
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += p.speed * elapsed
		hasMoved = true
		// facing |= gixel.Down
		// facing &= 0x1111 - gixel.Up
	}

	if !hasMoved {
		p.Animation.Play("idle", true)
	} else {
		p.Animation.Play("walk", false)
	}

	// p.SetFacing(facing)
	return nil
}
