package gixel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/graphic"
)

type BaseGxlSprite struct {
	BaseGxlObject
	graphic *graphic.GxlGraphic
}

// NewSprite creates a new instance of GxlSprite in a given position.
func NewSprite(x, y float64) GxlSprite {
	s := &BaseGxlSprite{}
	s.SetPosition(x, y)
	return s
}

// MakeGraphic creates a new GlxGraphic instance in form of a rectangle
// with given color and sets it as the sprite's graphic.
func (s *BaseGxlSprite) MakeGraphic(w, h int, c color.Color) {
	s.graphic = graphic.MakeGraphic(w, h, c)
	s.SetSize(w, h)
}

// LoadGraphic creates a new GlxGraphic from a file path
// and sets it as the sprite's graphic.
func (s *BaseGxlSprite) LoadGraphic(path string) {
	s.graphic = graphic.LoadGraphic(path)
	w, h := s.graphic.GetImage().Size()
	s.SetSize(w, h)
}

func (s *BaseGxlSprite) Draw(screen *ebiten.Image) {
	s.BaseGxlObject.Draw(screen)
	if s.graphic == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.Scale.X*s.FacingMult.X, s.Scale.Y*s.FacingMult.Y)
	op.GeoM.Rotate(s.Angle)
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.graphic.GetImage(), op)
}

type GxlSprite interface {
	GxlObject
	MakeGraphic(w, h int, c color.Color)
	LoadGraphic(path string)
}
