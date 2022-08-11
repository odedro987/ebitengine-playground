package gixel

import (
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type BaseGxlText struct {
	BaseGxlObject
	text     string
	color    color.Color
	img      *ebiten.Image
	tt       *sfnt.Font
	fontSize float64
}

func (t *BaseGxlText) Init() {
	t.BaseGxlObject.Init()
	if t.tt == nil {
		log.Fatal("cannot set nil font")
	}
	t.updateGraphic()
}

// NewText creates a new instance of GxlText with a given font in a given position.
func NewText(x, y float64, text string, font *sfnt.Font) GxlText {
	t := &BaseGxlText{}
	t.SetPosition(x, y)
	t.tt = font
	t.fontSize = 8
	t.color = color.White
	t.text = text
	return t
}

func (t *BaseGxlText) SetText(text string) {
	t.text = text
	t.updateGraphic()
}

func (t *BaseGxlText) SetFont(font *sfnt.Font) {
	if font == nil {
		log.Fatal("cannot set nil font")
	}
	t.tt = font
	t.updateGraphic()
}

func (t *BaseGxlText) SetFontSize(size float64) {
	t.fontSize = size
	t.updateGraphic()
}

func (t *BaseGxlText) SetColor(color color.Color) {
	t.color = color
	t.updateGraphic()
}

func (t *BaseGxlText) updateGraphic() {
	newFace, err := opentype.NewFace(t.tt, &opentype.FaceOptions{
		Size:    t.fontSize,
		DPI:     72, // TODO: Support high dpi displays
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatal(err)
	}

	rect := text.BoundString(newFace, t.text)
	p := rect.Size()
	t.w, t.h = p.X, p.Y

	t.img = ebiten.NewImage(t.w, t.h)
	text.Draw(t.img, t.text, newFace, -rect.Min.X, -rect.Min.Y, t.color)
}

func (t *BaseGxlText) Draw(screen *ebiten.Image) {
	t.BaseGxlObject.Draw(screen)
	if t.img == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(-t.w/2), float64(-t.h/2))
	op.GeoM.Rotate(t.Angle * t.FacingMult.X)
	op.GeoM.Scale(t.Scale.X*t.FacingMult.X, t.Scale.Y*t.FacingMult.Y)
	op.GeoM.Translate(float64(t.w/2), float64(t.h/2))
	op.GeoM.Translate(t.X, t.Y)

	screen.DrawImage(t.img, op)
}

type GxlText interface {
	GxlObject
	SetText(text string)
	SetFont(font *sfnt.Font)
	SetFontSize(size float64)
	SetColor(color color.Color)
	updateGraphic()
}
