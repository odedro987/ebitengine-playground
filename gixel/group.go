package gixel

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type BaseGxlGroup[T GxlBasic] struct {
	BaseGxlBasic
	members       []*T
	maxSize       int
	recycleMarker func() int
}

func NewGroup[T GxlBasic](maxSize int) *BaseGxlGroup[T] {
	if maxSize < 0 {
		log.Panicln("maxSize cannot be negative")
	}

	return &BaseGxlGroup[T]{
		maxSize: maxSize,
	}
}

func (g *BaseGxlGroup[T]) Init() {
	g.BaseGxlBasic.Init()
	if g.maxSize == 0 {
		g.members = make([]*T, 0)
	} else {
		g.members = make([]*T, 0, g.maxSize)
	}

	g.recycleMarker = cyclicCounter(0, g.maxSize-1)
}

func (g *BaseGxlGroup[T]) Add(object T) *T {
	freeSlotIdx := -1
	for idx, member := range g.members {
		if member == &object {
			log.Println("object already exists in group")
			return nil
		}

		if freeSlotIdx == -1 && member == nil {
			freeSlotIdx = idx
		}
	}

	if freeSlotIdx != -1 {
		g.members[freeSlotIdx] = &object
		object.Init()
		return &object
	}

	if g.maxSize > 0 && len(g.members) >= g.maxSize {
		log.Println("group limit exceeded")
		return nil
	}

	g.members = append(g.members, &object)
	object.Init()
	return &object
}

func (g *BaseGxlGroup[T]) getFirstAvailable() *T {
	for _, member := range g.members {
		if member != nil && !*(*member).Exists() {
			return member
		}
	}

	return nil
}

func (g *BaseGxlGroup[T]) Recycle(factory func() T) *T {
	// Case group has limit
	if g.maxSize > 0 {
		if len(g.members) < g.maxSize {
			obj := factory()
			g.Add(obj)
			return &obj
		} else {
			// On each recycle returns a ref to a member in a cyclic order
			return g.members[g.recycleMarker()]
		}
	} else { // Case group has no limit
		first := g.getFirstAvailable()

		if first != nil {
			return first
		}

		obj := factory()
		g.Add(obj)
		return &obj
	}
}

func (g *BaseGxlGroup[T]) Remove(object *T) *T {
	for idx, member := range g.members {
		if member == object {
			g.members[idx] = nil
			return object
		}
	}

	return nil
}

func (g *BaseGxlGroup[T]) Length() int {
	return len(g.members)
}

func (g *BaseGxlGroup[T]) Range(f func(idx int, value *T) bool) {
	for idx, m := range g.members {
		if m == nil {
			continue
		}

		if !f(idx, m) {
			break
		}
	}
}

func (g *BaseGxlGroup[T]) Draw(screen *ebiten.Image) {
	for _, m := range g.members {
		if m != nil && *(*m).Exists() && *(*m).Visible() {
			(*m).Draw(screen)
		}
	}
}

func (g *BaseGxlGroup[T]) Update(elapsed float64) error {
	for _, m := range g.members {
		if m != nil && *(*m).Exists() {
			err := (*m).Update(elapsed)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *BaseGxlGroup[T]) Destroy() {
	for _, m := range g.members {
		if m != nil {
			(*m).Destroy()
		}
	}
}

// TODO: Figure out if a package is needed
// cyclicCounter increments a value on func call, cycles between min and max.
func cyclicCounter(min, max int) func() int {
	i := min - 1
	return func() int {
		if i >= max {
			i = min - 1
		}
		i++
		return i
	}
}

type GxlGroup[T GxlBasic] interface {
	GxlBasic
	Add(object T) *T
	getFirstAvailable() *T
	Length() int
	Range(f func(idx int, value *T) bool)
	Recycle(factory func() T) *T
	Remove(object *T) *T
}
