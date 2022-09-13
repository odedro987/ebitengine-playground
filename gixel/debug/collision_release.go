//go:build !debug

package debug

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/math"
)

type Collision struct{}

func (c *Collision) Update() {}

func (c *Collision) DrawBounds(screen *ebiten.Image, bounds math.GxlRectangle) {}
