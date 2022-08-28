package main

import (
	"math/rand"
	"time"

	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/font"
)

type PlayState struct {
	gixel.BaseGxlState
	letters    gixel.GxlGroup
	spawnTimer *gixel.GxlTimer
}

func (s *PlayState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	s.letters = gixel.NewGroup(0)
	s.Add(s.letters)

	f := font.NewFont("assets/Boku2-Bold.otf")

	var katakana []rune

	alphabet := `アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ`

	for _, letter := range alphabet {
		katakana = append(katakana, letter)
	}

	rand := rand.New(rand.NewSource(time.Now().Unix()))
	s.spawnTimer = gixel.NewLoopTimer(1.0 / 60).SetCallback(func(totalElapsed float64, iteration int) {
		for i := 0; i < 20; i++ {
			letter := s.letters.Recycle(func() gixel.GxlBasic {
				return NewLetter(0, 0, string(katakana[rand.Int()%(len(katakana)-1)]), f.GetPreset(32))
			}).(*Letter)

			*letter.X() = rand.Float64() * GAME_WIDTH
			*letter.Y() = rand.Float64() * GAME_HEIGHT
			letter.alpha = 255
			*letter.Exists() = true
		}
	})
	s.spawnTimer.Start()
}

func (s *PlayState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	s.spawnTimer.Update(elapsed)

	return nil
}
