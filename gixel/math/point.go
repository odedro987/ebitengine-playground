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
