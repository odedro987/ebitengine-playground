package math

type MigPoint struct {
	X, Y float64
}

func NewPoint(x, y float64) *MigPoint {
	return &MigPoint{
		X: x,
		Y: y,
	}
}