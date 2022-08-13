package gixel

import (
	"github.com/solarlune/resolv"
)

type CallbackFunc = func(obj1, obj2 *GxlObject)

type BaseGxlState struct {
	BaseGxlGroup
	Game         *GxlGame
	space        *resolv.Space
	spaceObjects map[GxlObject]*resolv.Object
}

func (s *BaseGxlState) Init() {
	s.BaseGxlGroup.Init()
	// TODO: Change cell size to tiles
	s.space = resolv.NewSpace(s.Game.width, s.Game.height, 16, 16)
	s.spaceObjects = make(map[GxlObject]*resolv.Object)
}

func (s *BaseGxlState) SetGame(game *GxlGame) {
	s.Game = game
}

func (s *BaseGxlState) alignObjectsInSpace(objs ...GxlObject) {
	for _, obj := range objs {
		_, ok := s.spaceObjects[obj]
		x, y := obj.GetPosition()
		w, h := obj.GetSize()
		if !ok {
			s.spaceObjects[obj] = resolv.NewObject(x, y, float64(w), float64(h))
			s.space.Add(s.spaceObjects[obj])
			continue
		}

		if *obj.Immovable() {
			continue
		}

		s.spaceObjects[obj].X, s.spaceObjects[obj].Y = x, y
		s.spaceObjects[obj].W, s.spaceObjects[obj].H = float64(w), float64(h)
		s.spaceObjects[obj].Update()
	}
}

func (s *BaseGxlState) OverlapsObjects(obj1, obj2 GxlObject) bool {
	s.alignObjectsInSpace(obj1, obj2)
	return s.spaceObjects[obj1].Overlaps(s.spaceObjects[obj2])
}

func (s *BaseGxlState) OverlapsObjectGroup(obj GxlObject, grp GxlGroup, callbacks ...CallbackFunc) bool {
	overlapped := false

	grp.Range(func(idx int, value *GxlBasic) bool {
		cobj, ok := (*value).(GxlObject)
		if !ok {
			return true
		}

		if !s.OverlapsObjects(obj, cobj) {
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

type GxlState interface {
	GxlGroup
	SetGame(game *GxlGame)
	OverlapsObjects(obj1, obj2 GxlObject) bool
	OverlapsObjectGroup(obj GxlObject, grp GxlGroup, callbacks ...CallbackFunc) bool
	OverlapsGroups(grp1, grp2 GxlGroup, callbacks ...CallbackFunc) bool
}
