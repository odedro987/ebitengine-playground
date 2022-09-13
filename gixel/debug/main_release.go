//go:build !debug

package debug

import "github.com/hajimehoshi/ebiten/v2"

type Debug struct {
	Counters  Counters
	Collision Collision
}

type Counters struct {
	InitCount   Counter
	UpdateCount Counter
	DrawCount   Counter
}

func (dc *Counters) DrawInfo(screen *ebiten.Image) {}

type Counter bool // changed to bool to save memory

func (c *Counter) Increment() {}

func (c *Counter) Clear() {}
