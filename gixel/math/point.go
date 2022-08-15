package math

import (
	"log"
	"math"
)

type GxlPoint struct {
	X, Y float64
}

const TO_RAD = math.Pi / 180
const TO_DEG = 180 / math.Pi

// NewPoint creates a new instance of GxlPoint with given coordinates.
func NewPoint(x, y float64) *GxlPoint {
	return &GxlPoint{
		X: x,
		Y: y,
	}
}

// Set changes the point's coordinates to the given ones.
func (p *GxlPoint) Set(x, y float64) {
	p.X, p.Y = x, y
}

// Set copies the coordinates from a given point.
func (p *GxlPoint) Copy(p2 *GxlPoint) {
	p.X, p.Y = p2.X, p2.Y
}

// Mult multiplies the coordinates by the given point.
//
// Returns a new instance of a GxlPoint with the result.
func (p *GxlPoint) Mult(p2 *GxlPoint) GxlPoint {
	return GxlPoint{X: p.X * p2.X, Y: p.Y * p2.Y}
}

// Mult divides the coordinates by the given point.
//
// Returns a new instance of a GxlPoint with the result.
//
// **Panics on zero division**.
func (p *GxlPoint) Div(p2 *GxlPoint) GxlPoint {
	if p2.X == 0.0 || p2.Y == 0.0 {
		log.Panicln("zero division")
	}

	return GxlPoint{X: p.X / p2.X, Y: p.Y / p2.Y}
}

// Mult adds the coordinates by the given point.
//
// Returns a new instance of a GxlPoint with the result.
func (p *GxlPoint) Add(p2 *GxlPoint) GxlPoint {
	return GxlPoint{X: p.X + p2.X, Y: p.Y + p2.Y}
}

// Mult subtracts the coordinates by the given point.
//
// Returns a new instance of a GxlPoint with the result.
func (p *GxlPoint) Sub(p2 *GxlPoint) GxlPoint {
	return GxlPoint{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func (p *GxlPoint) SetRadians(rads float64) {
	length := p.Length()
	p.X = math.Cos(rads) * length
	p.Y = math.Sin(rads) * length
}

func (p *GxlPoint) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *GxlPoint) PivotRadians(pivot *GxlPoint, radians float64) GxlPoint {
	point1 := pivot.Sub(p)
	point1.SetRadians(radians)
	return GxlPoint{point1.X + pivot.X, point1.Y + pivot.Y}
}

func (p *GxlPoint) PivotDegrees(pivot *GxlPoint, degrees float64) GxlPoint {
	return p.PivotRadians(pivot, degrees*TO_RAD)
}
