package entities

import (
	"image/color"

	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/systems/collision"
	"github.com/odedro987/gixel-engine/gixel/systems/physics"
)

type Wall struct {
	gixel.BaseGxlSprite
	// Systems
	physics.Physics
	collision.Collision
}

func NewWall(x, y float64) *Wall {
	w := &Wall{}
	w.SetPosition(x, y)

	return w
}

func (w *Wall) Init(game *gixel.GxlGame) {
	w.BaseGxlSprite.Init(game)

	w.Physics.Init(w)
	w.Collision.Init(w)

	*w.Immovable() = true

	w.MakeGraphic(32, 32, color.RGBA{A: 255, B: 255})

}

func (w *Wall) Update(elapsed float64) error {
	err := w.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	w.Collision.Update(elapsed)
	w.Physics.Update(elapsed)

	return nil
}
