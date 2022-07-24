package game

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/pkg/state"
)

type GxlGame struct {
	width     int
	height    int
	title     string
	lastFrame time.Time
	state     state.GxlState
}

func NewGame(width, height int, title string, initialState state.GxlState) {
	g := GxlGame{
		width:     width,
		height:    height,
		title:     title,
		lastFrame: time.Now(),
	}
	g.SwitchState(initialState)

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
	defer func() { g.state.Destroy() }()

}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *GxlGame) Update() error {
	// Calculate time since last frame.
	elpased := time.Since(g.lastFrame).Seconds()
	defer func() { g.lastFrame = time.Now() }()

	err := g.state.Update(elpased)
	if err != nil {
		return err
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *GxlGame) Draw(screen *ebiten.Image) {
	g.state.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *GxlGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func (g *GxlGame) SwitchState(nextState state.GxlState) {
	if g.state != nil {
		g.state.Destroy()
	}

	g.state = nextState
	g.state.Init()
}
