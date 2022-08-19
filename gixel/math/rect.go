package math

type GxlRectangle struct {
	x, y, w, h float64
}

func NewRectangle(x, y, w, h float64) *GxlRectangle {
	return &GxlRectangle{x, y, w, h}
}

func (r *GxlRectangle) X() *float64 {
	return &r.x
}

func (r *GxlRectangle) Y() *float64 {
	return &r.y
}

func (r *GxlRectangle) W() *float64 {
	return &r.w
}

func (r *GxlRectangle) H() *float64 {
	return &r.h
}

func (r *GxlRectangle) Overlaps(r2 *GxlRectangle) bool {
	return r2.x < r.x+r.w && r2.x+r2.w > r.x && r2.y < r.y+r.h && r2.y+r2.h > r.y
}
