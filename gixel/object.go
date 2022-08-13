package gixel

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/math"
)

// TODO: Figure out a place to have this type
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
	x, y       float64
	w, h       int
	angle      float64
	scale      *math.GxlPoint
	facingFlip map[GxlDirection]math.GxlPoint
	facing     GxlDirection
	facingMult *math.GxlPoint
}

func (o *BaseGxlObject) Init() {
	o.BaseGxlBasic.Init()
	o.scale = math.NewPoint(1, 1)
	o.facingFlip = make(map[GxlDirection]math.GxlPoint)
	o.facingMult = math.NewPoint(1, 1)
	o.facing = None
}

func (o *BaseGxlObject) X() *float64 {
	return &o.x
}

func (o *BaseGxlObject) Y() *float64 {
	return &o.y
}

func (o *BaseGxlObject) GetPosition() (x, y float64) {
	return o.x, o.y
}

func (o *BaseGxlObject) SetPosition(x, y float64) {
	o.x, o.y = x, y
}

func (o *BaseGxlObject) GetSize() (w, h int) {
	return o.w, o.h
}

func (o *BaseGxlObject) SetSize(w, h int) {
	o.w, o.h = w, h
}

func (o *BaseGxlObject) Facing() *GxlDirection {
	return &o.facing
}

func (o *BaseGxlObject) Scale() *math.GxlPoint {
	return o.scale
}

func (o *BaseGxlObject) Angle() *float64 {
	return &o.angle
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
			s.facingMult.Copy(&flipMult)
		}
	}
}

type GxlObject interface {
	GxlBasic
	X() *float64
	Y() *float64
	GetPosition() (x, y float64)
	SetPosition(x, y float64)
	GetSize() (w, h int)
	SetSize(w, h int)
	Facing() *GxlDirection
	Scale() *math.GxlPoint
	Angle() *float64
	SetFacingFlip(dir GxlDirection, flipX, flipY bool)
}
