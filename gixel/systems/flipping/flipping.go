package flipping

import (
	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/math"
)

type flippingRequirements interface {
	ScaleMultiplier() *math.GxlPoint
	AngleMultiplier() *float64
}

type Flipping struct {
	subject    *flippingRequirements
	facingFlip map[gixel.GxlDirection]math.GxlPoint
	facing     gixel.GxlDirection
}

func (f *Flipping) Init(subject flippingRequirements) {
	f.subject = &subject
	f.facingFlip = make(map[gixel.GxlDirection]math.GxlPoint)
	f.facing = gixel.None
}

func (f *Flipping) Update() {
	if f.facing == gixel.None {
		return
	}

	flipMult, ok := f.facingFlip[f.facing]
	if !ok {
		return
	}

	*(*f.subject).AngleMultiplier() = flipMult.X

	(*f.subject).ScaleMultiplier().Copy(&flipMult)
}

func (f *Flipping) Facing() *gixel.GxlDirection {
	return &f.facing
}

func (f *Flipping) SetFacingFlip(dir gixel.GxlDirection, flipX, flipY bool) {
	x, y := 1.0, 1.0
	if flipX {
		x = -1
	}
	if flipY {
		y = -1
	}

	f.facingFlip[dir] = *math.NewPoint(x, y)
}
