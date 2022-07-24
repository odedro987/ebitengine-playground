package state

import (
	"github.com/odedro987/gixel-engine/pkg/basic"

	"github.com/hajimehoshi/ebiten/v2"
)

type Base struct {
	members []basic.GxlBasic
}

func (s *Base) Init() {
	s.members = make([]basic.GxlBasic, 0)
}

func (s *Base) Destroy() {
	for _, m := range s.members {
		m.Destroy()
	}
}

func (s *Base) Draw(screen *ebiten.Image) {
	for _, m := range s.members {
		if m.Exists() && m.IsVisible() {
			m.Draw(screen)
		}
	}
}

func (s *Base) Update(elapsed float64) error {
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

func (s *Base) Add(basic basic.GxlBasic) {
	basic.Init()
	s.members = append(s.members, basic)
}

type GxlState interface {
	Init()
	Destroy()
	Draw(screen *ebiten.Image)
	Update(elapsed float64) error
	Add(basic basic.GxlBasic)
}
