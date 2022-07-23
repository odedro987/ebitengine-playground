package state

import (
	"github.com/odedro987/ebitengine-playground/pkg/basic"

	"github.com/hajimehoshi/ebiten/v2"
)

type MigStateBase struct {
	members []basic.MigBasic
}

func (s *MigStateBase) Init() {
	s.members = make([]basic.MigBasic, 0)
}

func (s *MigStateBase) Destroy() {
	for _, m := range s.members {
		m.Destroy()
	}
}

func (s *MigStateBase) Draw(screen *ebiten.Image) {
	for _, m := range s.members {
		if m.Exists() && m.IsVisible() {
			m.Draw(screen)
		}
	}
}

func (s *MigStateBase) Update(elapsed float64) error {
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

func (s *MigStateBase) Add(basic basic.MigBasic) {
	basic.Init()
	s.members = append(s.members, basic)
}

type MigState interface {
	Init()
	Destroy()
	Draw(screen *ebiten.Image)
	Update(elapsed float64) error
	Add(basic basic.MigBasic)
}