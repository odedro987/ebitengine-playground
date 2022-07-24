package gixel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/animation"
	"github.com/odedro987/gixel-engine/gixel/graphic"
)

type BaseGxlSprite struct {
	BaseGxlObject
	graphic   *graphic.GxlGraphic
	Animation *animation.GxlAnimationController
}

func (s *BaseGxlSprite) Init() {
	s.BaseGxlObject.Init()
	s.Animation = animation.NewAnimationController()
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

// LoadGraphic creates a new GlxGraphic from a file path
// and sets it as the sprite's graphic.
func (s *BaseGxlSprite) LoadAnimatedGraphic(path string, fw, fh int) {
	s.graphic = graphic.LoadAnimatedGraphic(path, fw, fh)
	s.SetSize(fw, fh)
}

func (s *BaseGxlSprite) Update(elapsed float64) error {
	err := s.BaseGxlObject.Update(elapsed)
	if err != nil {
		return err
	}

	s.Animation.Update(elapsed)

	return nil
}

func (s *BaseGxlSprite) Draw(screen *ebiten.Image) {
	s.BaseGxlObject.Draw(screen)
	if s.graphic == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	w, h := s.graphic.GetSize()
	op.GeoM.Translate(float64(-w/2), float64(-h/2))
	op.GeoM.Rotate(s.Angle * s.FacingMult.X)
	op.GeoM.Scale(s.Scale.X*s.FacingMult.X, s.Scale.Y*s.FacingMult.Y)
	op.GeoM.Translate(float64(w/2), float64(h/2))
	op.GeoM.Translate(s.X, s.Y)

	frameIdx := 0
	if s.Animation.CurrAnim != nil {
		frameIdx = s.Animation.FrameIndex
	}

	screen.DrawImage(s.graphic.GetFrame(frameIdx), op)
}

type GxlSprite interface {
	GxlObject
	MakeGraphic(w, h int, c color.Color)
	LoadGraphic(path string)
}
