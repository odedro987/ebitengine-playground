package animation

import (
	"math"
)

type GxlAnimation struct {
	name          string
	isFinished    bool
	isLoop        bool
	isPaused      bool
	frameIndex    int
	maxFrameIndex int
	frameRate     float64
	timer         float64
	delay         float64
	frames        []int
	controller    *GxlAnimationController
}

func NewAnimation(controller *GxlAnimationController, name string, frames []int, fps float64, looped bool) *GxlAnimation {
	a := GxlAnimation{
		name:          name,
		frames:        frames,
		timer:         0,
		frameIndex:    0,
		maxFrameIndex: len(frames) - 1,
		isLoop:        looped,
		isFinished:    false,
		isPaused:      false,
		controller:    controller,
	}
	a.SetFPS(fps)

	return &a
}

func (a *GxlAnimation) SetFPS(fps float64) {
	a.frameRate = fps
	a.delay = 0
	if math.Abs(fps) > 0 {
		a.delay = 1 / math.Abs(fps)
	}
}

func (a *GxlAnimation) Pause() {
	a.isPaused = true
}

func (a *GxlAnimation) Resume() {
	a.isPaused = false
}

func (a *GxlAnimation) Stop() {
	a.isFinished = true
	a.isPaused = true
}

func (a *GxlAnimation) Reset() {
	a.Stop()
	a.frameIndex = 0
}

func (a *GxlAnimation) Restart() {
	a.Play(true)
}

func (a *GxlAnimation) Play(force bool) {
	if !force && !a.isFinished {
		a.isPaused = false
		// a.isFinished = false ??? Why flixel??
		return
	}

	a.isPaused = false
	a.isFinished = a.delay == 0
	a.frameIndex = 0
	a.timer = 0

	// if a.isFinished {
	// 	// OnFinish() ?
	// }
}

func (a *GxlAnimation) Update(elapsed float64) {
	if a.isFinished || a.delay == 0 || a.isPaused {
		a.controller.FrameIndex = a.frames[a.frameIndex]
		return
	}

	a.timer += elapsed

	if a.timer > a.delay && !a.isFinished {
		a.timer -= a.delay

		if a.frameIndex == a.maxFrameIndex {
			if a.isLoop {
				a.frameIndex = 0
			} else {
				a.isFinished = true
			}
		} else {
			a.frameIndex++
		}
	}

	a.controller.FrameIndex = a.frames[a.frameIndex]
}
