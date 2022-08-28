package animation

import (
	"log"

	"github.com/odedro987/gixel-engine/gixel/graphic"
)

type imports interface {
	Graphic() *graphic.GxlGraphic
	W() *int
	H() *int
}

type Animation struct {
	subject    *imports
	animations map[string]*GxlAnimation
	currAnim   *GxlAnimation
	frameIndex int
}

type Exports interface {
	AddAnimation(name string, frames []int, fps float64, looped bool) *GxlAnimation
	SetAnimationFPS(fps float64)
	PauseAnimation()
	ResumeAnimation()
	StopAnimation()
	ResetAnimation()
	RestartAnimation()
	PlayAnimation(name string, force bool)
}

func (a *Animation) Init(subject imports) {
	a.subject = &subject
	a.animations = make(map[string]*GxlAnimation)
}

func (a *Animation) AddAnimation(name string, frames []int, fps float64, looped bool) *GxlAnimation {
	a.animations[name] = NewAnimation(name, frames, fps, looped)
	return a.animations[name]
}

func (a *Animation) SetAnimationFPS(fps float64) {
	if a.currAnim != nil {
		a.currAnim.setFPS(fps)
	}
}

func (a *Animation) PauseAnimation() {
	if a.currAnim != nil {
		a.currAnim.pause()
	}
}

func (a *Animation) ResumeAnimation() {
	if a.currAnim != nil {
		a.currAnim.resume()
	}
}

func (a *Animation) StopAnimation() {
	if a.currAnim != nil {
		a.currAnim.stop()
	}
}

func (a *Animation) ResetAnimation() {
	if a.currAnim != nil {
		a.currAnim.reset()
	}
}

func (a *Animation) RestartAnimation() {
	if a.currAnim != nil {
		a.currAnim.restart()
	}
}

func (a *Animation) PlayAnimation(name string, force bool) {
	anim, ok := a.animations[name]
	if !ok {
		log.Printf("no animation called %s\n", name)
		return
	}

	if a.currAnim != nil && a.currAnim.getName() != name {
		a.currAnim.reset()
	}

	a.currAnim = anim
	a.currAnim.play(force)
}

func (a *Animation) Update(elapsed float64) {
	if a.currAnim == nil || (*a.subject).Graphic() == nil {
		return
	}

	a.currAnim.update(elapsed)

	a.frameIndex = a.currAnim.getCurrentFrame()
	(*a.subject).Graphic().SetCurrentFrameIdx(a.frameIndex)
}
