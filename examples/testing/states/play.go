package states

import (
	"github.com/odedro987/mig-engine/examples/testing/entities"
	"github.com/odedro987/mig-engine/pkg/sprite"
	"github.com/odedro987/mig-engine/pkg/state"
)

type PlayState struct {
	state.Base
	player sprite.GxlSprite
}

func (s *PlayState) Init() {
	s.Base.Init()
	s.player = entities.NewPlayer(100, 100)
	s.Add(s.player)
}