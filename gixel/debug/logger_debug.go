//go:build debug

package debug

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type logEntry struct {
	Message   string
	Timestamp time.Time
}

type GxlLogger struct {
	logs []*logEntry
	ttl  time.Duration
}

func NewLogger(ttl time.Duration) *GxlLogger {
	logger := GxlLogger{
		ttl: ttl,
	}
	log.SetFlags(0) // Remove timestamp prefix
	log.SetOutput(&logger)
	return &logger
}

func (l *GxlLogger) getFirstAvailable() *logEntry {
	for i := len(l.logs) - 1; i >= 0; i-- {
		if l.logs[i] == nil {
			return l.logs[i]
		}
	}

	return nil
}

func (l *GxlLogger) Write(p []byte) (n int, err error) {
	msg := string(p)

	log := logEntry{
		Message:   msg,
		Timestamp: time.Now(),
	}
	first := l.getFirstAvailable()
	if first != nil {
		*first = log
	} else {
		l.logs = append(l.logs, &log)
	}

	fmt.Fprint(os.Stderr, msg)
	return len(p), nil
}

func (l *GxlLogger) Draw(screen *ebiten.Image) {
	output := ""
	for i := len(l.logs) - 1; i >= 0; i-- {
		if l.logs[i] != nil {
			if time.Since(l.logs[i].Timestamp) >= l.ttl {
				l.logs[i] = nil
				continue
			}
			output += l.logs[i].Message
		}
	}
	ebitenutil.DebugPrint(screen, output)
}
