package states

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel"
	"golang.org/x/image/font/opentype"
)

type MenuState struct {
	gixel.BaseGxlState
}

func (s *MenuState) Init() {
	s.BaseGxlState.Init()

	box1 := gixel.NewSprite(100, 100)
	box1.MakeGraphic(100, 100, color.RGBA{R: 255, A: 255})

	box2 := gixel.NewSprite(200, 200)
	box2.MakeGraphic(100, 100, color.RGBA{B: 255, A: 255})

	fontBytes, err := os.ReadFile("assets/nokiafc22.ttf")
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	text := gixel.NewText(100, 100, "Hello World", tt)

	s.Add(box1)
	s.Add(box2)
	s.Add(text)
}

func (s *MenuState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	if ebiten.IsKeyPressed(ebiten.KeyP) {
		s.Game.SwitchState(&PlayState{})
	}

	return nil
}
