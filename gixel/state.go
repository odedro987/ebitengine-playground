package gixel

type BaseGxlState struct {
	BaseGxlGroup[GxlBasic]
	Game *GxlGame
}

func (s *BaseGxlState) SetGame(game *GxlGame) {
	s.Game = game
}

type GxlState interface {
	GxlGroup[GxlBasic]
	SetGame(game *GxlGame)
}
