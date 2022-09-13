package graphic

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type GxlGraphic struct {
	frames []*ebiten.Image
}

func (g *GxlGraphic) GetFrame(idx int) *ebiten.Image {
	if idx < 0 || idx >= len(g.frames) {
		log.Panicln("index out of bounds")
	}
	return g.frames[idx]
}

func (g *GxlGraphic) Size() (int, int) {
	return g.frames[0].Size()
}

func (g *GxlGraphic) FrameWidth() int {
	w, _ := g.frames[0].Size()
	return w
}

func (g *GxlGraphic) FrameHeight() int {
	_, h := g.frames[0].Size()
	return h
}
