package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type screens int32

type typeCard struct {
	CardValue int32
	CardTeam  int32
}
type Card struct {
	TypeCard  typeCard
	RecSource rl.Rectangle
}

const (
	Clubs = iota
	Diamonds
	Hearts
	Spades
)

// card valuoes
const (
	As = 1
	J  = 11
	Q  = 12
	K  = 13
)

// screens
const (
	Start screens = iota
	Client
	Server
)

type game struct {
	RayLogo     rayLogo
	Cards       [4][13]Card
	GameScreen  screens
	Spritesheet rl.Texture2D
	Map         map[typeCard]rl.Rectangle
	Hand        []typeCard
}
type rayLogo struct {
	logoPositionX      int32
	logoPositionY      int32
	framesCounter      int32
	lettersCount       int32
	topSideRecWidth    int32
	leftSideRecHeight  int32
	bottomSideRecWidth int32
	rightSideRecHeight int32
	state              int32
	alpha              float32
}

func initSprites() map[typeCard]rl.Rectangle {
	Map := make(map[typeCard]rl.Rectangle)
	for y := 0; y < 4; y++ {
		for x := 0; x < 13; x++ {
			switch y {
			case int(Clubs):
				Map[typeCard{CardValue: int32(x + 1), CardTeam: Clubs}] = rl.Rectangle{
					X:      float32(x * 29),
					Y:      0,
					Width:  29,
					Height: 36,
				}
			case int(Diamonds):
				Map[typeCard{CardValue: int32(x + 1), CardTeam: Diamonds}] = rl.Rectangle{
					X:      float32((x * 29) + 376),
					Y:      0,
					Width:  29,
					Height: 36,
				}

			case int(Hearts):
				Map[typeCard{CardValue: int32(x + 1), CardTeam: Hearts}] = rl.Rectangle{
					X:      float32((x * 29) + 753),
					Y:      0,
					Width:  29,
					Height: 36,
				}
			case int(Spades):
				Map[typeCard{CardValue: int32(x + 1), CardTeam: Spades}] = rl.Rectangle{
					X:      float32((x * 29) + 1130),
					Y:      0,
					Width:  29,
					Height: 36,
				}
			}
		}
	}
	return Map
}
