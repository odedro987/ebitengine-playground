package states

import (
	"image/color"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/odedro987/gixel-engine/gixel"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"golang.org/x/image/font/opentype"
)

type MenuState struct {
	gixel.BaseGxlState
	text              *gixel.GxlText
	textScaleSequence *gween.Sequence
	textAngleSequence *gween.Sequence
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

	text.SetFontSize(16)
	s.text = &text

	s.textScaleSequence = gween.NewSequence(
		gween.New(1, 2, 5, ease.InQuad),
		gween.New(2, 1, 5, ease.InQuad),
	)
	s.textScaleSequence.SetLoop(-1)

	s.textAngleSequence = gween.NewSequence(
		gween.New(-math.Pi*0.25, math.Pi*0.25, 2, ease.InOutElastic),
		gween.New(math.Pi*0.25, -math.Pi*0.25, 2, ease.InOutElastic),
	)
	s.textAngleSequence.SetLoop(-1)

	s.Add(box1)
	s.Add(box2)
	s.Add(text)
}

func (s *MenuState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	currentScale, _, _ := s.textScaleSequence.Update(float32(elapsed))
	(*s.text).Scale().X = float64(currentScale)
	(*s.text).Scale().Y = float64(currentScale)

	currentAngle, _, _ := s.textAngleSequence.Update(float32(elapsed))
	*(*s.text).Angle() = float64(currentAngle)

	if ebiten.IsKeyPressed(ebiten.KeyP) {
		s.Game.SwitchState(&PlayState{})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		fullscreen := ebiten.IsFullscreen()
		fullscreen = !fullscreen
		ebiten.SetFullscreen(fullscreen)
	}

	return nil
}
