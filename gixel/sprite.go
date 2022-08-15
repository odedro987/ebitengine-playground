package gixel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/graphic"
)

type BaseGxlSprite struct {
	BaseGxlObject
	img   *ebiten.Image
	color color.RGBA // TODO: Think of a better name
}

func (s *BaseGxlSprite) Init(game *GxlGame) {
	s.BaseGxlObject.Init(game)
	s.color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
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
	s.img = graphic.MakeGraphic(w, h, c).GetFrame(0)
	s.SetSize(w, h)
}

// LoadGraphic creates a new GlxGraphic from a file path
// and sets it as the sprite's graphic.
func (s *BaseGxlSprite) LoadGraphic(path string) {
	s.img = graphic.LoadGraphic(path).GetFrame(0)
	w, h := s.img.Size()
	s.SetSize(w, h)
}

func (s *BaseGxlSprite) Image() **ebiten.Image {
	return &s.img
}

func (s *BaseGxlSprite) Color() *color.RGBA {
	return &s.color
}

func (s *BaseGxlSprite) Update(elapsed float64) error {
	err := s.BaseGxlObject.Update(elapsed)
	if err != nil {
		return err
	}

	return nil
}

func (s *BaseGxlSprite) Draw(screen *ebiten.Image) {
	s.BaseGxlObject.Draw(screen)
	if s.img == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	w, h := s.img.Size()
	op.GeoM.Translate(float64(-w/2), float64(-h/2))
	op.GeoM.Rotate(s.angle * s.angleMultiplier)
	op.GeoM.Scale(s.scale.X*s.scaleMultiplier.X, s.scale.Y*s.scaleMultiplier.Y)
	op.GeoM.Translate(float64(w/2), float64(h/2))
	op.GeoM.Translate(s.x, s.y)
	// TODO: Add color for tinting/etc
	op.ColorM.ScaleWithColor(s.color)

	screen.DrawImage(s.img, op)
}

type GxlSprite interface {
	GxlObject
	MakeGraphic(w, h int, c color.Color)
	LoadGraphic(path string)
	Image() **ebiten.Image
	Color() *color.RGBA
}
