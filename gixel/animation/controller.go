package animation

import "fmt"

type GxlAnimationController struct {
	animations map[string]*GxlAnimation
	CurrAnim   *GxlAnimation
	FrameIndex int
}

func NewAnimationController() *GxlAnimationController {
	ac := GxlAnimationController{
		CurrAnim: nil,
	}
	ac.animations = make(map[string]*GxlAnimation)

	return &ac
}

func (ac *GxlAnimationController) Add(name string, frames []int, fps float64, looped bool) {
	ac.animations[name] = NewAnimation(ac, name, frames, fps, looped)
}

func (ac *GxlAnimationController) SetFPS(fps float64) {
	if ac.CurrAnim != nil {
		ac.CurrAnim.SetFPS(fps)
	}
}

func (ac *GxlAnimationController) Pause() {
	if ac.CurrAnim != nil {
		ac.CurrAnim.Pause()
	}
}

func (ac *GxlAnimationController) Resume() {
	if ac.CurrAnim != nil {
		ac.CurrAnim.Resume()
	}
}

func (ac *GxlAnimationController) Stop() {
	if ac.CurrAnim != nil {
		ac.CurrAnim.Stop()
	}
}

func (ac *GxlAnimationController) Reset() {
	if ac.CurrAnim != nil {
		ac.CurrAnim.Reset()
	}
}

func (ac *GxlAnimationController) Restart() {
	if ac.CurrAnim != nil {
		ac.CurrAnim.Restart()
	}
}

func (ac *GxlAnimationController) Play(name string, force bool) {
	anim, ok := ac.animations[name]
	if !ok {
		fmt.Printf("No animation called %s", name)
		return
	}

	if ac.CurrAnim != nil && ac.CurrAnim.name != name {
		ac.CurrAnim.Reset()
	}

	ac.CurrAnim = anim
	ac.CurrAnim.Play(force)
}

func (s *GxlAnimationController) Update(elapsed float64) {
	if s.CurrAnim != nil {
		s.CurrAnim.Update(elapsed)
	}
}
