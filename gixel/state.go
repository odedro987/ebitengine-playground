package gixel

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BaseGxlState struct {
	Game    *GxlGame
	members []GxlBasic
}

func (s *BaseGxlState) Init() {
	s.members = make([]GxlBasic, 0)
}

func (s *BaseGxlState) Destroy() {
	for _, m := range s.members {
		m.Destroy()
	}
}

func (s *BaseGxlState) Draw(screen *ebiten.Image) {
	for _, m := range s.members {
		if m.Exists() && m.IsVisible() {
			m.Draw(screen)
		}
	}
}

func (s *BaseGxlState) Update(elapsed float64) error {
	for _, m := range s.members {
		if m.Exists() {
			err := m.Update(elapsed)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *BaseGxlState) Add(basic GxlBasic) {
	basic.Init()
	s.members = append(s.members, basic)
}

type GxlState interface {
	Init()
	Destroy()
	Draw(screen *ebiten.Image)
	Update(elapsed float64) error
	Add(basic GxlBasic)
}
