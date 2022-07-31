package math

import "log"

type GxlPoint struct {
	X, Y float64
}

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
