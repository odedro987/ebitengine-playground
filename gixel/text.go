package gixel

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/odedro987/gixel-engine/gixel/font"
)

type BaseGxlText struct {
	BaseGxlSprite
	text       string
	fontPreset *font.GxlFontPreset
}

func (t *BaseGxlText) Init(game *GxlGame) {
	t.BaseGxlObject.Init(game)
	t.updateGraphic()
}

// NewText creates a new instance of GxlText with a given font in a given position.
func NewText(x, y float64, text string, fontPreset *font.GxlFontPreset) GxlText {
	t := &BaseGxlText{}
	t.SetPosition(x, y)
	t.fontPreset = fontPreset
	t.color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	t.text = text
	return t
}

func (t *BaseGxlText) SetText(text string) {
	t.text = text
	t.updateGraphic()
}

func (t *BaseGxlText) SetFontPreset(fontPreset *font.GxlFontPreset) {
	if fontPreset == nil {
		log.Fatal("cannot set nil font")
	}
	t.fontPreset = fontPreset
	t.updateGraphic()
}

func (t *BaseGxlText) Color() *color.RGBA {
	return &t.color
}

func (t *BaseGxlText) updateGraphic() {
	rect := text.BoundString(t.fontPreset.GetFace(), t.text)
	p := rect.Size()
	t.w, t.h = p.X, p.Y

	t.img = ebiten.NewImage(t.w, t.h)
	text.Draw(t.img, t.text, t.fontPreset.GetFace(), -rect.Min.X, -rect.Min.Y, color.White)
}

type GxlText interface {
	GxlSprite
	SetText(text string)
	SetFontPreset(fontPreset *font.GxlFontPreset)
	Color() *color.RGBA
	updateGraphic()
}
