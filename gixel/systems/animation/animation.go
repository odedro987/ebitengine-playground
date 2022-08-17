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
	callbacks     map[int]func()
	onFinished    func()
}

func NewAnimation(name string, frames []int, fps float64, looped bool) *GxlAnimation {
	a := GxlAnimation{
		name:          name,
		frames:        frames,
		timer:         0,
		frameIndex:    0,
		maxFrameIndex: len(frames) - 1,
		isLoop:        looped,
		isFinished:    true,
		isPaused:      true,
		callbacks:     make(map[int]func()),
	}
	a.setFPS(fps)

	return &a
}

func (a *GxlAnimation) SetCallback(frameIndex int, callback func()) *GxlAnimation {
	a.callbacks[frameIndex] = callback
	return a
}

func (a *GxlAnimation) SetOnFinished(callback func()) *GxlAnimation {
	a.onFinished = callback
	return a
}

func (a *GxlAnimation) setFPS(fps float64) {
	a.frameRate = fps
	a.delay = 0
	if math.Abs(fps) > 0 {
		a.delay = 1 / math.Abs(fps)
	}
}

func (a *GxlAnimation) pause() {
	a.isPaused = true
}

func (a *GxlAnimation) resume() {
	a.isPaused = false
}

func (a *GxlAnimation) stop() {
	a.isFinished = true
	a.isPaused = true
}

func (a *GxlAnimation) reset() {
	a.stop()
	a.frameIndex = 0
}

func (a *GxlAnimation) restart() {
	a.play(true)
}

func (a *GxlAnimation) getCurrentFrame() int {
	return a.frames[a.frameIndex]
}

func (a *GxlAnimation) getName() string {
	return a.name
}

func (a *GxlAnimation) play(force bool) {
	if !force && !a.isFinished {
		a.isPaused = false
		return
	}

	a.isPaused = false
	a.isFinished = a.delay == 0
	a.frameIndex = 0
	a.timer = 0

	callback, ok := a.callbacks[0]
	if ok {
		callback()
	}
}

func (a *GxlAnimation) update(elapsed float64) {
	if a.isFinished || a.delay == 0 || a.isPaused {
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
				if a.onFinished != nil {
					a.onFinished()
					return
				}
			}
		} else {
			a.frameIndex++
		}

		callback, ok := a.callbacks[a.frameIndex]
		if ok {
			callback()
		}
	}
}
