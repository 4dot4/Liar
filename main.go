package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Card struct {
	Number    cardValue
	TypeCard  typeCard
	RecSource rl.Rectangle
	RecDes    rl.Rectangle
}

//w 28
//h 35

func main() {
	var cards [4][13]Card

	for y := 0; y < 4; y++ {
		for x := 0; x < 13; x++ {
			switch y {
			case int(Clubs):
				cards[y][x].TypeCard = Clubs
				cards[y][x].Number = cardValue(x + 1)
				cards[y][x].RecSource = rl.Rectangle{
					X:      float32(x * 29),
					Y:      0,
					Width:  28,
					Height: 36,
				}
			case int(Diamonds):
				cards[y][x].TypeCard = Diamonds
				cards[y][x].RecSource = rl.Rectangle{
					X:      float32((x * 29) + 376),
					Y:      0,
					Width:  28,
					Height: 36,
				}
			case int(Hearts):
				cards[y][x].TypeCard = Hearts
				cards[y][x].RecSource = rl.Rectangle{
					X:      float32((x * 29) + 753),
					Y:      0,
					Width:  28,
					Height: 36,
				}
			case int(Spades):
				cards[y][x].TypeCard = Spades
				cards[y][x].RecSource = rl.Rectangle{
					X:      float32((x*28 + 1) + 1130),
					Y:      0,
					Width:  28,
					Height: 36,
				}

			}
			cards[y][x].Number = cardValue(x + 1)
		}
	}

	rl.InitWindow(800, 800, "Liar")
	var Spritesheet rl.Texture2D = rl.LoadTexture("./assets/spritesheet.png")
	x := 0

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyD) {
			if x < 3 {
				x++
			}
		}
		if rl.IsKeyPressed(rl.KeyA) {
			if x > 0 {
				x--
			}
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		for i := 0; i < 13; i++ {
			rl.DrawTextureRec(Spritesheet, cards[x][i].RecSource, rl.Vector2{X: float32(28 * i), Y: 1}, rl.White)
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
