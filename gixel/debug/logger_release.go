//go:build !debug

package debug

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type GxlLogger struct{}

func NewLogger(ttl time.Duration) *GxlLogger {
	return &GxlLogger{}
}

func (l *GxlLogger) Draw(screen *ebiten.Image) {}
