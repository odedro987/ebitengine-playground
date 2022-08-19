package gixel

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BaseGxlBasic struct {
	game    *GxlGame
	id      int
	visible bool
	exists  bool
	// TODO: Do we want current state ref here?
}

func (b *BaseGxlBasic) Init(g *GxlGame) {
	b.game = g
	b.id = g.GenerateID()
	b.visible = true
	b.exists = true
}

func (b *BaseGxlBasic) Destroy() {
	b.exists = false
}

func (b *BaseGxlBasic) Draw(screen *ebiten.Image) {}

func (b *BaseGxlBasic) Update(elapsed float64) error {
	return nil
}

func (b *BaseGxlBasic) GetID() int {
	return b.id
}

func (b *BaseGxlBasic) Visible() *bool {
	return &b.visible
}

func (b *BaseGxlBasic) Exists() *bool {
	return &b.exists
}

func (b *BaseGxlBasic) Game() *GxlGame {
	return b.game
}

type GxlBasic interface {
	Init(g *GxlGame)
	Destroy()
	Draw(screen *ebiten.Image)
	Update(elapsed float64) error
	GetID() int
	Visible() *bool
	Exists() *bool
	Game() *GxlGame
}
