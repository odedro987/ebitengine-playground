package gixel

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/math"
)

type GxlDirection uint16

const (
	None  GxlDirection = 0x0000
	Up    GxlDirection = 0x0001
	Down  GxlDirection = 0x0010
	Left  GxlDirection = 0x0100
	Right GxlDirection = 0x1000
)

type BaseGxlObject struct {
	BaseGxlBasic
	X, Y       float64
	w, h       int
	Angle      float64
	Scale      *math.GxlPoint
	facingFlip map[GxlDirection]math.GxlPoint
	facing     GxlDirection
	FacingMult *math.GxlPoint
}

func (o *BaseGxlObject) Init() {
	o.BaseGxlBasic.Init()
	o.Scale = math.NewPoint(1, 1)
	o.facingFlip = make(map[GxlDirection]math.GxlPoint)
	o.FacingMult = math.NewPoint(1, 1)
	o.facing = None
}

func (o *BaseGxlObject) GetPosition() (x, y float64) {
	return o.X, o.Y
}

func (o *BaseGxlObject) SetPosition(x, y float64) {
	o.X, o.Y = x, y
}

func (o *BaseGxlObject) GetSize() (w, h int) {
	return o.w, o.h
}

func (o *BaseGxlObject) SetSize(w, h int) {
	o.w, o.h = w, h
}

func (o *BaseGxlObject) GetFacing() GxlDirection {
	return o.facing
}

func (o *BaseGxlObject) SetFacing(dir GxlDirection) {
	o.facing = dir
}

func (o *BaseGxlObject) SetFacingFlip(dir GxlDirection, flipX, flipY bool) {
	x, y := 1.0, 1.0
	if flipX {
		x = -1
	}
	if flipY {
		y = -1
	}

	o.facingFlip[dir] = *math.NewPoint(x, y)
}

func (s *BaseGxlObject) Draw(screen *ebiten.Image) {
	s.BaseGxlBasic.Draw(screen)

	// Set `FacingMult` based on `facing`
	if s.facing != None {
		flipMult, ok := s.facingFlip[s.facing]
		if ok {
			s.FacingMult.Copy(&flipMult)
		}
	}
}

type GxlObject interface {
	GxlBasic
	GetPosition() (x, y float64)
	SetPosition(x, y float64)
	GetSize() (w, h int)
	SetSize(w, h int)
	GetFacing() GxlDirection
	SetFacing(dir GxlDirection)
	SetFacingFlip(dir GxlDirection, flipX, flipY bool)
}
