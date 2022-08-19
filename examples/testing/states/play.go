package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/odedro987/gixel-engine/examples/testing/entities"
	"github.com/odedro987/gixel-engine/gixel"
)

type PlayState struct {
	gixel.BaseGxlState
	player *entities.Player
	walls  *gixel.BaseGxlGroup
}

func (s *PlayState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	s.walls = gixel.NewGroup(0)
	s.Add(s.walls)
	for i := 0.0; i < 4; i++ {
		s.walls.Recycle(func() gixel.GxlBasic {
			return entities.NewWall(float64(100+32*i), 150+2*i)
		})
	}

	s.player = entities.NewPlayer(100, 100)
	s.Add(s.player)
}

func (s *PlayState) Draw(screen *ebiten.Image) {
	s.BaseGxlState.Draw(screen)
	// ebitenutil.DebugPrint(screen, strconv.Itoa(s.testGroup.Length()))
}

func (s *PlayState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	s.CollideObjectGroup(s.player, s.walls)

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
