package sprite

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/ebitengine-playground/pkg/graphic"
	"github.com/odedro987/ebitengine-playground/pkg/object"
)

type Base struct {
	object.Base
	graphic *graphic.MigGraphic
}

func New(x, y float64) (MigSprite) {
	s := &Base{}
	s.SetPosition(x, y)
	return s
}

func (s *Base) MakeGraphic(w, h int, c color.Color) {
	s.graphic = graphic.MakeGraphic(w, h, c)
	s.SetSize(w, h)
}

func (s *Base) Draw(screen *ebiten.Image) {
	if s.graphic == nil {
		return
	}
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.graphic.GetImage(), op)
}

type MigSprite interface {
	object.MigObject
	MakeGraphic(w, h int, c color.Color)
}