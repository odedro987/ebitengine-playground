//go:build !debug

package debug

import "github.com/hajimehoshi/ebiten/v2"

type DebugCounters struct {
	InitCount   Counter
	UpdateCount Counter
	DrawCount   Counter
}

func (dc *DebugCounters) DrawDebugInfo(screen *ebiten.Image) {}

type Counter bool // changed to bool to save memory

func (c *Counter) Increment() {}

func (c *Counter) Clear() {}
