package gixel

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	gl "github.com/odedro987/gixel-engine/gixel/log"
)

type GxlGame struct {
	width       int
	height      int
	title       string
	lastFrame   time.Time
	state       GxlState
	zoom        int
	logger      *gl.GxlLogger
	isDebug     bool // TODO: Expose
	timescale   float64
	nextID      int
	stateLoaded bool
}

func NewGame(width, height int, title string, initialState GxlState, zoom int) {
	g := GxlGame{
		width:       width,
		height:      height,
		title:       title,
		lastFrame:   time.Now(),
		zoom:        zoom,
		isDebug:     true,
		timescale:   1,
		nextID:      0,
		stateLoaded: false,
	}
	g.SwitchState(initialState)

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	// Use custom GxlLogger
	g.logger = gl.NewLogger(time.Second * 5)
	log.SetFlags(0) // Remove timestamp prefix
	log.SetOutput(g.logger)

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
	defer func() { g.state.Destroy() }()

}

func (g *GxlGame) GenerateID() int {
	id := g.nextID
	g.nextID++
	return id
}

func (g *GxlGame) W() int {
	return g.width
}

func (g *GxlGame) H() int {
	return g.height
}

// Update proceeds the game
// Update is called every tick (1/60 [s] by default).
func (g *GxlGame) Update() error {
	if !g.stateLoaded {
		return nil
	}

	// TODO: Figure out what to do with TPS
	elapsed := 1.0 / float64(ebiten.MaxTPS())

	err := g.state.Update(g.timescale * elapsed)
	if err != nil {
		return err
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *GxlGame) Draw(screen *ebiten.Image) {
	if !g.stateLoaded {
		return
	}

	g.state.Draw(screen)
	if g.isDebug {
		g.logger.Draw(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *GxlGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width / g.zoom, g.height / g.zoom
}

func (g *GxlGame) State() *GxlState {
	return &g.state
}

// SwitchState changes the game's current state and initializes it.
// It also calls the current state's Destroy func.
func (g *GxlGame) SwitchState(nextState GxlState) {
	if g.state != nil {
		g.state.Destroy()
	}
	g.stateLoaded = false
	g.state = nextState
	g.state.Init(g)
	g.stateLoaded = true
}

func (g *GxlGame) Timescale() *float64 {
	return &g.timescale
}
