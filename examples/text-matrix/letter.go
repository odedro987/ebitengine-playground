package main

import (
	"image/color"

	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/font"
)

type Letter struct {
	gixel.BaseGxlText
	alpha float64
}

func NewLetter(x float64, y float64, text string, fontPreset *font.GxlFontPreset) *Letter {
	t := &Letter{}
	t.SetPosition(x, y)
	t.SetText(text)
	t.SetFontPreset(fontPreset)
	return t
}

func (l *Letter) Init(game *gixel.GxlGame) {
	l.BaseGxlText.Init(game)
	*l.Color() = color.RGBA{G: 255, A: 255}
	l.alpha = 255
}

func (l *Letter) Update(elapsed float64) error {
	err := l.BaseGxlText.Update(elapsed)
	if err != nil {
		return err
	}

	l.alpha -= elapsed * 300

	l.Color().A = uint8(l.alpha)

	if l.alpha <= 0 {
		l.Destroy()
	}

	return nil
}
