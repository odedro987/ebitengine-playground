package animation

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/graphic"
)

type animationRequirements interface {
	Image() **ebiten.Image
	W() *int
	H() *int
}

type Animation struct {
	subject    *animationRequirements
	graphic    *graphic.GxlGraphic
	animations map[string]*GxlAnimation
	currAnim   *GxlAnimation
	frameIndex int
}

func (f *Animation) Init(subject animationRequirements) {
	f.subject = &subject
	f.animations = make(map[string]*GxlAnimation)
}

// LoadAnimatedGraphic creates a new GlxGraphic from a file path
// and sets it as the sprite's graphic.
func (f *Animation) LoadAnimatedGraphic(path string, fw, fh int) {
	f.graphic = graphic.LoadAnimatedGraphic(path, fw, fh)
	*(*f.subject).W() = fw
	*(*f.subject).H() = fh
}

func (f *Animation) AddAnimation(name string, frames []int, fps float64, looped bool) *GxlAnimation {
	f.animations[name] = NewAnimation(name, frames, fps, looped)
	return f.animations[name]
}

func (f *Animation) SetAnimationFPS(fps float64) {
	if f.currAnim != nil {
		f.currAnim.setFPS(fps)
	}
}

func (f *Animation) PauseAnimation() {
	if f.currAnim != nil {
		f.currAnim.pause()
	}
}

func (f *Animation) ResumeAnimation() {
	if f.currAnim != nil {
		f.currAnim.resume()
	}
}

func (f *Animation) StopAnimation() {
	if f.currAnim != nil {
		f.currAnim.stop()
	}
}

func (f *Animation) ResetAnimation() {
	if f.currAnim != nil {
		f.currAnim.reset()
	}
}

func (f *Animation) RestartAnimation() {
	if f.currAnim != nil {
		f.currAnim.restart()
	}
}

func (f *Animation) PlayAnimation(name string, force bool) {
	anim, ok := f.animations[name]
	if !ok {
		log.Printf("no animation called %s\n", name)
		return
	}

	if f.currAnim != nil && f.currAnim.getName() != name {
		f.currAnim.reset()
	}

	f.currAnim = anim
	f.currAnim.play(force)
}

func (s *Animation) Update(elapsed float64) {
	if s.currAnim == nil || s.graphic == nil {
		return
	}

	s.currAnim.update(elapsed)

	s.frameIndex = s.currAnim.getCurrentFrame()
	*(*s.subject).Image() = s.graphic.GetFrame(s.frameIndex)
}
