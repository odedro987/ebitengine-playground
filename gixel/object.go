package gixel

import (
	"github.com/odedro987/gixel-engine/gixel/math"
)

type BaseGxlObject struct {
	BaseGxlBasic
	X, Y  float64
	w, h  int
	Angle float64
	Scale *math.GxlPoint
}

func New(x, y float64, w, h int) (obj GxlObject) {
	return &BaseGxlObject{
		X: x,
		Y: y,
		w: w,
		h: h,
	}
}

func (o *BaseGxlObject) Init() {
	o.BaseGxlBasic.Init()
	o.Scale = math.NewPoint(1, 1)
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

type GxlObject interface {
	GxlBasic
	GetPosition() (x, y float64)
	SetPosition(x, y float64)
	GetSize() (w, h int)
	SetSize(w, h int)
}
