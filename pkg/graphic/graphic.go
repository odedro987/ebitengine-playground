package graphic

import (
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

func LoadGraphic(path string) (*MigGraphic) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
		// TODO: Error handling, default value?
	}
	
	return &MigGraphic{
		img: img,
	}
}

func (g *MigGraphic) GetImage() (*ebiten.Image) {
	return g.img
}