package basic

import "github.com/hajimehoshi/ebiten/v2"

type MigBasicBase struct {
	visible bool
	exists  bool
}

func (b *MigBasicBase) Init() {
	b.visible = true
	b.exists = true
}

func (b *MigBasicBase) Destroy() {
	b.exists = false
}

func (b *MigBasicBase) Draw(screen *ebiten.Image) {}

func (b *MigBasicBase) Update(elapsed float64) error {
	return nil
}

func (b *MigBasicBase) IsVisible() bool {
	return b.visible
}

func (b *MigBasicBase) Exists() bool {
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