package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type typeCard int32
type cardValue int32
type screens int32
type Card struct {
	Number    cardValue
	TypeCard  typeCard
	RecSource rl.Rectangle
	RecDes    rl.Rectangle
}

const (
	Clubs typeCard = iota
	Diamonds
	Hearts
	Spades
)

// card valuoes
const (
	As cardValue = 1
	J  cardValue = 11
	Q  cardValue = 12
	K  cardValue = 13
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

func initSprites(cards *[4][13]Card) {
	for y := 0; y < 4; y++ {
		for x := 0; x < 13; x++ {
			switch y {
			case int(Clubs):
				cards[y][x].TypeCard = Clubs
				cards[y][x].Number = cardValue(x + 1)
				cards[y][x].RecSource = rl.Rectangle{
					X:      float32(x * 29),
					Y:      0,
					Width:  29,
					Height: 36,
				}
			case int(Diamonds):
				cards[y][x].TypeCard = Diamonds
				cards[y][x].RecSource = rl.Rectangle{
					X:      float32((x * 29) + 376),
					Y:      0,
					Width:  29,
					Height: 36,
				}
			case int(Hearts):
				cards[y][x].TypeCard = Hearts
				cards[y][x].RecSource = rl.Rectangle{
					X:      float32((x * 29) + 753),
					Y:      0,
					Width:  29,
					Height: 36,
				}
			case int(Spades):
				cards[y][x].TypeCard = Spades
				cards[y][x].RecSource = rl.Rectangle{
					X:      float32((x * 29) + 1130),
					Y:      0,
					Width:  29,
					Height: 36,
				}

			}
			cards[y][x].Number = cardValue(x + 1)
		}
	}
}
