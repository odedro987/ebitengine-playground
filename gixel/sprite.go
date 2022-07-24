package gixel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/graphic"
	"github.com/odedro987/gixel-engine/gixel/math"
)

type BaseGxlSprite struct {
	BaseGxlObject
	graphic *graphic.GxlGraphic
}

func NewSprite(x, y float64) GxlSprite {
	s := &BaseGxlSprite{}
	s.SetPosition(x, y)
	return s
}

func (s *BaseGxlSprite) MakeGraphic(w, h int, c color.Color) {
	s.graphic = graphic.MakeGraphic(w, h, c)
	s.SetSize(w, h)
}

func (s *BaseGxlSprite) LoadGraphic(path string) {
	s.graphic = graphic.LoadGraphic(path)
	w, h := s.graphic.GetImage().Size()
	s.SetSize(w, h)
}

func (s *BaseGxlSprite) Draw(screen *ebiten.Image) {
	if s.graphic == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(s.Angle)
	op.GeoM.Translate(s.X, s.Y)
	if *s.Scale != *math.NewPoint(1, 1) {
		op.GeoM.Scale(s.Scale.X, s.Scale.Y)
	}
	screen.DrawImage(s.graphic.GetImage(), op)
}

type GxlSprite interface {
	GxlObject
	MakeGraphic(w, h int, c color.Color)
}
