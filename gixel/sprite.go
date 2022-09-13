package gixel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/graphic"
)

type BaseGxlSprite struct {
	BaseGxlObject
	graphic  *graphic.GxlGraphic
	frameIdx int
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

func (s *BaseGxlSprite) ApplyGraphic(graphic *graphic.GxlGraphic) {
	s.graphic = graphic
	s.SetSize(graphic.Size())
}

func (s *BaseGxlSprite) Graphic() *graphic.GxlGraphic {
	return s.graphic
}

func (s *BaseGxlSprite) FrameIdx() *int {
	return &s.frameIdx
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

func (s *BaseGxlSprite) Draw() {
	if !s.OnScreen() {
		return
	}

	s.BaseGxlObject.Draw()
	if s.graphic == nil {
		return
	}

	s.drawOpts.GeoM.Reset()
	sx, sy := s.ScreenPosition()
	w, h := s.graphic.Size()
	s.drawOpts.GeoM.Translate(float64(-w/2), float64(-h/2))
	s.drawOpts.GeoM.Rotate(s.angle * s.angleMultiplier)
	s.drawOpts.GeoM.Scale(s.scale.X*s.scaleMultiplier.X, s.scale.Y*s.scaleMultiplier.Y)
	s.drawOpts.GeoM.Translate(float64(w/2), float64(h/2))
	s.drawOpts.GeoM.Translate(sx, sy)
	// // TODO: Add color for tinting/etc
	s.drawOpts.ColorM.Reset()
	s.drawOpts.ColorM.ScaleWithColor(s.color)

	s.camera.Screen().DrawImage(s.graphic.GetFrame(s.frameIdx), s.drawOpts)
}

type GxlSprite interface {
	GxlObject
	// MakeGraphic(w, h int, c color.Color)
	// LoadGraphic(path string)
	// LoadAnimatedGraphic(path string, fw, fh int)
	ApplyGraphic(graphic *graphic.GxlGraphic)
	Graphic() *graphic.GxlGraphic
	Color() *color.RGBA
}
