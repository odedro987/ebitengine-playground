package states

import (
	"image/color"

	"github.com/odedro987/gixel-engine/gixel"
)

type MenuState struct {
	gixel.BaseGxlState
}

func (s *MenuState) Init() {
	s.BaseGxlState.Init()

	box1 := gixel.NewSprite(100, 100)
	box1.MakeGraphic(100, 100, color.RGBA{R: 255, A: 255})

	box2 := gixel.NewSprite(100, 100)
	box2.MakeGraphic(100, 100, color.RGBA{B: 255, A: 255})

	s.Add(box1)
	s.Add(box2)
}
