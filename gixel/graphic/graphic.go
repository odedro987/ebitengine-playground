package graphic

import (
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GxlGraphic struct {
	frames []*ebiten.Image
}

// MakeGraphic creates a new ebiten.Image fills it with a given color.
//
// Returns a pointer of the GxlGraphic.
func MakeGraphic(w, h int, c color.Color) *GxlGraphic {
	img := ebiten.NewImage(w, h)
	img.Fill(c)

	return &GxlGraphic{
		frames: []*ebiten.Image{img},
	}
}

// MakeGraphic creates a new ebiten.Image from a file path.
//
// Returns a pointer of the GxlGraphic.
func LoadGraphic(path string) *GxlGraphic {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Panicln(err)
		// TODO: Error handling, default value?
	}

	return &GxlGraphic{
		frames: []*ebiten.Image{img},
	}
}

// MakeGraphic creates a new ebiten.Image from a file path.
//
// Returns a pointer of the GxlGraphic.
func LoadAnimatedGraphic(path string, fw, fh int) *GxlGraphic {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Panicln(err)
		// TODO: Error handling, default value?
	}

	w, h := img.Size()
	frameRows := (h / fh)
	frameCols := (w / fw)

	frames := make([]*ebiten.Image, 0, frameCols*frameRows)

	for i := 0; i < frameRows; i++ {
		for j := 0; j < frameCols; j++ {
			frameRect := image.Rect(j*fw, i*fh, j*fw+fw, i*fh+fh)
			frames = append(frames, img.SubImage(frameRect).(*ebiten.Image))
		}
	}

	return &GxlGraphic{
		frames: frames,
	}
}

func (g *GxlGraphic) GetFrames() *[]*ebiten.Image {
	return &g.frames
}

func (g *GxlGraphic) GetFrame(idx int) *ebiten.Image {
	if idx < 0 || idx >= len(g.frames) {
		log.Panicln("index out of bounds")
	}
	return g.frames[idx]
}

func (g *GxlGraphic) GetSize() (int, int) {
	return g.frames[0].Size()
}
