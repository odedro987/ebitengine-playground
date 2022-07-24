package math

type GxlPoint struct {
	X, Y float64
}

func NewPoint(x, y float64) *GxlPoint {
	return &GxlPoint{
		X: x,
		Y: y,
	}
}

func (p *GxlPoint) Set(x, y float64) {
	p.X, p.Y = x, y
}

func (p *GxlPoint) Copy(p2 *GxlPoint) {
	p.X, p.Y = p2.X, p2.Y
}

func (p *GxlPoint) Mult(p2 *GxlPoint) GxlPoint {
	return GxlPoint{X: p.X * p2.X, Y: p.Y * p2.Y}
}

func (p *GxlPoint) Div(p2 *GxlPoint) GxlPoint {
	if p2.X == 0.0 || p2.Y == 0.0 {
		panic("Zero division")
	}

	return GxlPoint{X: p.X / p2.X, Y: p.Y / p2.Y}
}

func (p *GxlPoint) Add(p2 *GxlPoint) GxlPoint {
	return GxlPoint{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p *GxlPoint) Sub(p2 *GxlPoint) GxlPoint {
	return GxlPoint{X: p.X - p2.X, Y: p.Y - p2.Y}
}
