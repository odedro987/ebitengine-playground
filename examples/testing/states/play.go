package states

import (
	"image/color"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/odedro987/gixel-engine/gixel"
)

type PlayState struct {
	gixel.BaseGxlState
	testGroup *gixel.BaseGxlGroup[gixel.GxlSprite]
}

func (s *PlayState) Init() {
	s.BaseGxlState.Init()

	s.testGroup = gixel.NewGroup[gixel.GxlSprite](0)

	s.Add(s.testGroup)
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

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		obj := s.testGroup.Recycle(func() gixel.GxlSprite {
			return gixel.NewSprite(0, 0)
		})
		(*obj).SetPosition(0, float64(rand.Int()%300))
		(*obj).MakeGraphic(50, 50, color.White)
	}

	s.testGroup.Range(
		func(_ int, member *gixel.GxlSprite) bool {
			posX, posY := (*member).GetPosition()
			(*member).SetPosition(posX+elapsed*40, posY)
			// if rand.Float64() > 0.99 {
			// 	s.testGroup.Remove(member)
			// }
			return true
		},
	)

	return nil
}
