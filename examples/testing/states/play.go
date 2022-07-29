package states

import (
	"image/color"

	"github.com/odedro987/gixel-engine/examples/testing/entities"
	"github.com/odedro987/gixel-engine/gixel"
)

type PlayState struct {
	gixel.BaseGxlState
	testGroup *gixel.GxlGroup[gixel.GxlSprite]
}

func (s *PlayState) Init() {
	s.BaseGxlState.Init()

	s.testGroup = gixel.NewGroup[gixel.GxlSprite](0)

	player1 := entities.NewPlayer(100, 100)
	player2 := entities.NewPlayer(100, 200)
	testSprite := gixel.NewSprite(50, 50)
	testSprite.MakeGraphic(50, 50, color.White)

	s.testGroup.Add(player1)
	s.testGroup.Add(player2)

	s.Add(s.testGroup)
}

func (s *PlayState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	s.testGroup.Range(
		func(_ int, member *gixel.GxlSprite) bool {
			posX, posY := (*member).GetPosition()
			(*member).SetPosition(posX+elapsed*10, posY)
			return true
		},
	)

	return nil
}
