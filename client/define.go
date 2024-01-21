package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)



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
