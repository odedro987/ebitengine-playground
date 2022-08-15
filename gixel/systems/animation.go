package systems

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/gixel-engine/gixel/animation"
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
	animations map[string]*animation.GxlAnimation
	currAnim   *animation.GxlAnimation
	frameIndex int
}

func (f *Animation) Init(subject animationRequirements) {
	f.subject = &subject
	f.animations = make(map[string]*animation.GxlAnimation)
}

// LoadAnimatedGraphic creates a new GlxGraphic from a file path
// and sets it as the sprite's graphic.
func (f *Animation) LoadAnimatedGraphic(path string, fw, fh int) {
	f.graphic = graphic.LoadAnimatedGraphic(path, fw, fh)
	*(*f.subject).W() = fw
	*(*f.subject).H() = fh
}

func (f *Animation) AddAnimation(name string, frames []int, fps float64, looped bool) {
	f.animations[name] = animation.NewAnimation(name, frames, fps, looped)
}

func (f *Animation) SetAnimationFPS(fps float64) {
	if f.currAnim != nil {
		f.currAnim.SetFPS(fps)
	}
}

func (f *Animation) PauseAnimation() {
	if f.currAnim != nil {
		f.currAnim.Pause()
	}
}

func (f *Animation) ResumeAnimation() {
	if f.currAnim != nil {
		f.currAnim.Resume()
	}
}

func (f *Animation) StopAnimation() {
	if f.currAnim != nil {
		f.currAnim.Stop()
	}
}

func (f *Animation) ResetAnimation() {
	if f.currAnim != nil {
		f.currAnim.Reset()
	}
}

func (f *Animation) RestartAnimation() {
	if f.currAnim != nil {
		f.currAnim.Restart()
	}
}

func (f *Animation) PlayAnimation(name string, force bool) {
	anim, ok := f.animations[name]
	if !ok {
		log.Printf("no animation called %s\n", name)
		return
	}

	if f.currAnim != nil && f.currAnim.GetName() != name {
		f.currAnim.Reset()
	}

	f.currAnim = anim
	f.currAnim.Play(force)
}

func (s *Animation) Update(elapsed float64) {
	if s.currAnim == nil || s.graphic == nil {
		return
	}

	s.currAnim.Update(elapsed)

	s.frameIndex = s.currAnim.GetCurrentFrame()
	*(*s.subject).Image() = s.graphic.GetFrame(s.frameIndex)
}
