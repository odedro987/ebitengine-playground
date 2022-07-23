package basic

import "github.com/hajimehoshi/ebiten/v2"

type Base struct {
	visible bool
	exists  bool
}

func (b *Base) Init() {
	b.visible = true
	b.exists = true
}

func (b *Base) Destroy() {
	b.exists = false
}

func (b *Base) Draw(screen *ebiten.Image) {}

func (b *Base) Update(elapsed float64) error {
	return nil
}

func (b *Base) IsVisible() bool {
	return b.visible
}

func (b *Base) Exists() bool {
	return b.exists
}

type MigBasic interface {
	Init()
	Destroy()
	Draw(screen *ebiten.Image)
	Update(elapsed float64) error
	IsVisible() bool
	Exists() bool
}
