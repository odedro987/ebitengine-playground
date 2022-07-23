package graphic

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type MigGraphic struct {
	img *ebiten.Image
}

func MakeGraphic(w, h int, c color.Color) (*MigGraphic) {
	img := ebiten.NewImage(w, h)
	img.Fill(c)
	
	return &MigGraphic{
		img: img,
	}
}

func (g *MigGraphic) GetImage() (*ebiten.Image) {
	return g.img
}