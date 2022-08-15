package states

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/odedro987/gixel-engine/examples/testing/entities"
	"github.com/odedro987/gixel-engine/gixel"
)

type PlayState struct {
	gixel.BaseGxlState
	player     *entities.Player
	testGroup  *gixel.BaseGxlGroup
	testGroup2 *gixel.BaseGxlGroup
}

func (s *PlayState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	s.testGroup = gixel.NewGroup(0)
	s.testGroup2 = gixel.NewGroup(0)

	s.player = entities.NewPlayer(100, 100)

	b1 := gixel.NewSprite(150, 150)
	b1.MakeGraphic(50, 50, color.RGBA{R: 255, A: 255})
	b2 := gixel.NewSprite(50, 50)
	b2.MakeGraphic(50, 50, color.RGBA{G: 255, A: 255})

	s.Add(s.testGroup)
	s.Add(s.testGroup2)
	s.testGroup2.Add(s.player)
	s.testGroup.Add(b1)
	s.testGroup.Add(b2)
}

func (s *PlayState) Draw(screen *ebiten.Image) {
	s.BaseGxlState.Draw(screen)
	ebitenutil.DebugPrint(screen, strconv.Itoa(s.testGroup.Length()))
}

func (s *PlayState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	s.OverlapsGroups(
		s.testGroup2,
		s.testGroup,
		func(obj1, obj2 *gixel.GxlObject) {
			*(*obj2).X() += 20 * elapsed
		},
	)

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		fullscreen := ebiten.IsFullscreen()
		fullscreen = !fullscreen
		ebiten.SetFullscreen(fullscreen)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		*s.Game().Timescale() = 0.5
	}

	return nil
}
