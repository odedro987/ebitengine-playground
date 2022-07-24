package object

import (
	"github.com/odedro987/mig-engine/pkg/basic"
	"github.com/odedro987/mig-engine/pkg/math"
)

type Base struct {
	basic.Base
	X, Y  float64
	w, h  int
	Angle float64
	Scale *math.GxlPoint
}

func New(x, y float64, w, h int) (obj GxlObject) {
	return &Base{
		X: x,
		Y: y,
		w: w,
		h: h,
	}
}

func (o *Base) Init() {
	o.Base.Init()
	o.Scale = math.NewPoint(1, 1)
}

func (o *Base) GetPosition() (x, y float64) {
	return o.X, o.Y
}

func (o *Base) SetPosition(x, y float64) {
	o.X, o.Y = x, y
}

func (o *Base) GetSize() (w, h int) {
	return o.w, o.h
}

func (o *Base) SetSize(w, h int) {
	o.w, o.h = w, h
}

type GxlObject interface {
	basic.GxlBasic
	GetPosition() (x, y float64)
	SetPosition(x, y float64)
	GetSize() (w, h int)
	SetSize(w, h int)
}
