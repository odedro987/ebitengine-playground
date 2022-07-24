package graphic

import (
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GxlGraphic struct {
	img *ebiten.Image
}

// MakeGraphic creates a new ebiten.Image fills it with a given color.
//
// Returns a pointer of the GxlGraphic.
func MakeGraphic(w, h int, c color.Color) *GxlGraphic {
	img := ebiten.NewImage(w, h)
	img.Fill(c)

	return &GxlGraphic{
		img: img,
	}
}

// MakeGraphic creates a new ebiten.Image from a file path.
//
// Returns a pointer of the GxlGraphic.
func LoadGraphic(path string) *GxlGraphic {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
		// TODO: Error handling, default value?
	}

	return &GxlGraphic{
		img: img,
	}
}

func (g *GxlGraphic) GetImage() *ebiten.Image {
	return g.img
}
