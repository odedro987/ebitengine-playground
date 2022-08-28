package gixel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/graphic"
)

type BaseGxlSprite struct {
	BaseGxlObject
	graphic  *graphic.GxlGraphic
	color    color.RGBA // TODO: Think of a better name
	drawOpts *ebiten.DrawImageOptions
}

func (s *BaseGxlSprite) Init(game *GxlGame) {
	s.BaseGxlObject.Init(game)
	s.drawOpts = &ebiten.DrawImageOptions{}
	s.color = color.RGBA{255, 255, 255, 255}
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
	w, h := s.graphic.GetSize()
	s.SetSize(w, h)
}

// LoadAnimatedGraphic creates a new GlxGraphic from a file path
// and sets it as the sprite'a graphic.
func (s *BaseGxlSprite) LoadAnimatedGraphic(path string, fw, fh int) {
	s.graphic = graphic.LoadAnimatedGraphic(path, fw, fh)
	s.w, s.h = fw, fh
}

func (s *BaseGxlSprite) Graphic() *graphic.GxlGraphic {
	return s.graphic
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
	if s.graphic == nil {
		return
	}

	s.drawOpts.GeoM.Reset()
	w, h := s.graphic.GetSize()
	s.drawOpts.GeoM.Translate(float64(-w/2), float64(-h/2))
	s.drawOpts.GeoM.Rotate(s.angle * s.angleMultiplier)
	s.drawOpts.GeoM.Scale(s.scale.X*s.scaleMultiplier.X, s.scale.Y*s.scaleMultiplier.Y)
	s.drawOpts.GeoM.Translate(float64(w/2), float64(h/2))
	s.drawOpts.GeoM.Translate(s.x, s.y)
	// // TODO: Add color for tinting/etc
	s.drawOpts.ColorM.Reset()
	s.drawOpts.ColorM.ScaleWithColor(s.color)

	screen.DrawImage(s.graphic.GetCurrentFrame(), s.drawOpts)
}

type GxlSprite interface {
	GxlObject
	MakeGraphic(w, h int, c color.Color)
	LoadGraphic(path string)
	LoadAnimatedGraphic(path string, fw, fh int)
	Graphic() *graphic.GxlGraphic
	Color() *color.RGBA
}
