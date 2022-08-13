package gixel

import "github.com/hajimehoshi/ebiten/v2"

type BaseGxlBasic struct {
	visible bool
	exists  bool
}

func (b *BaseGxlBasic) Init() {
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

func (b *BaseGxlBasic) Visible() *bool {
	return &b.visible
}

func (b *BaseGxlBasic) Exists() *bool {
	return &b.exists
}

type GxlBasic interface {
	Init()
	Destroy()
	Draw(screen *ebiten.Image)
	Update(elapsed float64) error
	Visible() *bool
	Exists() *bool
}
