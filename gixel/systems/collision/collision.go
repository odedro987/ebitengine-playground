package collision

import (
	"github.com/odedro987/gixel-engine/gixel/math"
)

type imports interface {
	Velocity() *math.GxlPoint
	Acceleration() *math.GxlPoint
	X() *float64
	Y() *float64
}

type Exports interface {
	LastX() float64
	LastY() float64
	Immovable() *bool
}

type Collision struct {
	subject   *imports
	immovable bool
	currPos   *math.GxlPoint
	lastPos   *math.GxlPoint
}

func (c *Collision) Init(subject imports) {
	c.subject = &subject
	c.lastPos = math.NewPoint(0, 0)
	c.currPos = math.NewPoint(0, 0)
}

func (c *Collision) Immovable() *bool {
	return &c.immovable
}

func (c *Collision) LastX() float64 {
	return c.lastPos.X
}

func (c *Collision) LastY() float64 {
	return c.lastPos.Y
}

func (c *Collision) Update(elapsed float64) {
	c.lastPos.Copy(c.currPos)

	x, y := *(*c.subject).X(), *(*c.subject).Y()
	c.currPos.Set(x, y)
}
