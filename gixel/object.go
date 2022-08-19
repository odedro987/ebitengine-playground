package gixel

import (
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
	x, y            float64
	w, h            int
	angle           float64
	angleMultiplier float64
	scale           *math.GxlPoint
	scaleMultiplier *math.GxlPoint
}

func (o *BaseGxlObject) Init(game *GxlGame) {
	o.BaseGxlBasic.Init(game)
	o.scale = math.NewPoint(1, 1)
	o.scaleMultiplier = math.NewPoint(1, 1)
	o.angleMultiplier = 1
}

func (o *BaseGxlObject) X() *float64 {
	return &o.x
}

func (o *BaseGxlObject) Y() *float64 {
	return &o.y
}

func (o *BaseGxlObject) W() *int {
	return &o.w
}

func (o *BaseGxlObject) H() *int {
	return &o.h
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

func (o *BaseGxlObject) Scale() *math.GxlPoint {
	return o.scale
}

func (o *BaseGxlObject) ScaleMultiplier() *math.GxlPoint {
	return o.scaleMultiplier
}

func (o *BaseGxlObject) Angle() *float64 {
	return &o.angle
}

func (o *BaseGxlObject) AngleMultiplier() *float64 {
	return &o.angleMultiplier
}

func (o *BaseGxlObject) Bounds() *math.GxlRectangle {
	return math.NewRectangle(o.x, o.y, float64(o.w), float64(o.h))
}

func (o *BaseGxlObject) Overlaps(obj GxlObject) bool {
	if !o.exists || !*obj.Exists() {
		return false
	}

	return o.Bounds().Overlaps(obj.Bounds())
}

type GxlObject interface {
	GxlBasic
	X() *float64
	Y() *float64
	W() *int
	H() *int
	GetPosition() (x, y float64)
	SetPosition(x, y float64)
	GetSize() (w, h int)
	SetSize(w, h int)
	Scale() *math.GxlPoint
	ScaleMultiplier() *math.GxlPoint
	Angle() *float64
	AngleMultiplier() *float64
	Bounds() *math.GxlRectangle
	Overlaps(obj GxlObject) bool
}
