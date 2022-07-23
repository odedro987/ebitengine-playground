package states

import (
	"image/color"

	"github.com/odedro987/mig-engine/pkg/sprite"
	"github.com/odedro987/mig-engine/pkg/state"
)

type MenuState struct {
	state.Base
}

func (s *MenuState) Init() {
	s.Base.Init()

	box1 := sprite.New(100, 100)
	box1.MakeGraphic(100, 100, color.RGBA{R: 255, A: 255})

	box2 := sprite.New(100, 100)
	box2.MakeGraphic(100, 100, color.RGBA{B: 255, A: 255})

	s.Add(box1)
	s.Add(box2)
}
