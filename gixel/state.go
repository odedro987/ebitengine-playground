package gixel

import (
	"math"

	"github.com/odedro987/gixel-engine/gixel/systems/collision"
)

type CallbackFunc = func(obj1, obj2 *GxlObject)

type BaseGxlState struct {
	BaseGxlGroup
}

type CollidableObject interface {
	GxlObject
	collision.Exports
}

// TODO: Figure out small delta collisions (1 pixel difference)
func (s *BaseGxlState) CollideObjects(obj1, obj2 CollidableObject) {
	if !obj1.Overlaps(obj2) {
		return
	}

	overlapX := computeOverlapX(obj1, obj2)
	overlapY := computeOverlapY(obj1, obj2)

	if (!*obj1.Immovable() && !*obj2.Immovable()) || *obj1.Immovable() {
		*obj2.X() -= overlapX
		*obj2.Y() -= overlapY
	} else if *obj2.Immovable() {
		*obj1.X() -= overlapX
		*obj1.Y() -= overlapY
	}
}

func (s *BaseGxlState) CollideObjectGroup(obj CollidableObject, grp GxlGroup) {
	grp.Range(func(idx int, value *GxlBasic) bool {
		cobj, ok := (*value).(CollidableObject)
		if !ok {
			return true
		}

		s.CollideObjects(obj, cobj)
		return true
	})
}

func (s *BaseGxlState) CollideGroups(grp1, grp2 GxlGroup) {
	grp1.Range(func(idx int, value *GxlBasic) bool {
		cobj, ok := (*value).(CollidableObject)
		if !ok {
			return true
		}

		s.OverlapsObjectGroup(cobj, grp2)
		return true
	})
}

func (s *BaseGxlState) OverlapsObjectGroup(obj GxlObject, grp GxlGroup, callbacks ...CallbackFunc) bool {
	overlapped := false

	grp.Range(func(idx int, value *GxlBasic) bool {
		cobj, ok := (*value).(GxlObject)
		if !ok {
			return true
		}

		if !obj.Overlaps(cobj) {
			return true
		}

		overlapped = true
		for _, callback := range callbacks {
			callback(&obj, &cobj)
		}
		return true
	})

	return overlapped
}

func (s *BaseGxlState) OverlapsGroups(grp1, grp2 GxlGroup, callbacks ...CallbackFunc) bool {
	overlapped := false

	grp1.Range(func(idx int, value *GxlBasic) bool {
		cobj, ok := (*value).(GxlObject)
		if !ok {
			return true
		}

		overlapped = overlapped || s.OverlapsObjectGroup(cobj, grp2, callbacks...)
		return true
	})

	return overlapped
}

func computeOverlapX(obj1, obj2 CollidableObject) float64 {
	// calculate how much overlap
	delta1 := *obj1.X() - obj1.LastX()
	delta2 := *obj2.X() - obj2.LastX()

	if delta1 == delta2 {
		return 0
	}

	delta1Abs := math.Abs(delta1)
	delta2Abs := math.Abs(delta2)
	maxOverlap := delta1Abs + delta2Abs + 4
	overlap := 0.0

	// TODO: Figure out if necessary
	// rect1 := gm.NewRectangle(*obj1.X()-math.Max(0.0, delta1), obj1.LastY(), float64(*obj1.W())+delta1Abs, float64(*obj1.H()))
	// rect2 := gm.NewRectangle(*obj2.X()-math.Max(0.0, delta2), obj2.LastY(), float64(*obj2.W())+delta2Abs, float64(*obj2.H()))

	// if !rect1.Overlaps(rect2) {
	// 	return 0
	// }

	if delta1 > delta2 {
		overlap = *obj1.X() + float64(*obj1.W()) - *obj2.X()
		if overlap > maxOverlap {
			return 0
		}
	} else if delta1 < delta2 {
		overlap = *obj1.X() - float64(*obj2.W()) - *obj2.X()
		if -overlap > maxOverlap {
			return 0
		}
	}

	return overlap
}

func computeOverlapY(obj1, obj2 CollidableObject) float64 {
	// calculate how much overlap
	delta1 := *obj1.Y() - obj1.LastY()
	delta2 := *obj2.Y() - obj2.LastY()

	if delta1 == delta2 {
		return 0
	}

	delta1Abs := math.Abs(delta1)
	delta2Abs := math.Abs(delta2)
	maxOverlap := delta1Abs + delta2Abs + 4
	overlap := 0.0

	if delta1 > delta2 {
		overlap = *obj1.Y() + float64(*obj1.H()) - *obj2.Y()
		if overlap > maxOverlap {
			return 0
		}
	} else if delta1 < delta2 {
		overlap = *obj1.Y() - float64(*obj2.H()) - *obj2.Y()
		if -overlap > maxOverlap {
			return 0
		}
	}

	return overlap
}

type GxlState interface {
	GxlGroup
	OverlapsObjectGroup(obj GxlObject, grp GxlGroup, callbacks ...CallbackFunc) bool
	OverlapsGroups(grp1, grp2 GxlGroup, callbacks ...CallbackFunc) bool
}
