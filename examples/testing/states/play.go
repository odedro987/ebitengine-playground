package states

import (
	"github.com/odedro987/gixel-engine/examples/testing/entities"
	"github.com/odedro987/gixel-engine/gixel"
)

type PlayState struct {
	gixel.BaseGxlState
	player gixel.GxlSprite
}

func (s *PlayState) Init() {
	s.BaseGxlState.Init()
	s.player = entities.NewPlayer(100, 100)
	s.Add(s.player)
}
