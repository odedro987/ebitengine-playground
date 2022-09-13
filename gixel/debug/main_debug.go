//go:build debug

package debug

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Debug struct {
	Counters  Counters
	Collision Collision
}

type Counters struct {
	InitCount   Counter
	UpdateCount Counter
	DrawCount   Counter
}

func (dc *Counters) DrawInfo(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%f - TPS:%f || I:%d - U:%d - D:%d", ebiten.ActualFPS(), ebiten.ActualTPS(), dc.InitCount, dc.UpdateCount, dc.DrawCount))
}

type Counter uint64

func (c *Counter) Increment() {
	*c++
}

func (c *Counter) Clear() {
	*c = 0
}
