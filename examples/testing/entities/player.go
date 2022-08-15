package entities

import (
	"github.com/odedro987/gixel-engine/examples/testing/systems"
	"github.com/odedro987/gixel-engine/gixel"
	gs "github.com/odedro987/gixel-engine/gixel/systems"
)

type Player struct {
	gixel.BaseGxlSprite
	gs.Physics
	systems.Movement
}

func NewPlayer(x, y float64) *Player {
	p := &Player{}
	p.SetPosition(x, y)

	return p
}

func (p *Player) Init() {
	p.BaseGxlSprite.Init()
	p.LoadAnimatedGraphic("assets/player.png", 32, 32)
	p.SetFacingFlip(gixel.Right, false, false)
	p.SetFacingFlip(gixel.Left, true, false)

	p.Animation().Add("WalkFront", []int{0, 1, 0, 2}, 7, true)
	p.Animation().Add("WalkBack", []int{3, 4, 3, 5}, 7, true)
	p.Animation().Add("WalkSide", []int{6, 7, 6, 8}, 7, true)

	p.Animation().Add("StandFront", []int{0, 0, 9}, 5, true)
	p.Animation().Add("StandBack", []int{3, 3, 10}, 5, true)

	p.Physics.Init(p)
	p.Movement.Init(p)
}

func (p *Player) Update(elapsed float64) error {
	err := p.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	p.Physics.Update(elapsed)
	p.Movement.Update(elapsed)
	p.Animation().Play(p.Movement.GetAnimName(), false)
	return nil
}
