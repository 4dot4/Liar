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

func main() {
	go server()
	var cards [4][13]Card
	var Width int32 = 800
	var Height int32 = 800
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

	rl.InitWindow(Width, Height, "Liar")
	rl.SetTargetFPS(60)

	var Spritesheet rl.Texture2D = rl.LoadTexture("../assets/spritesheet.png")
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

			rl.DrawTexturePro(Spritesheet, cards[x][i].RecSource, rl.Rectangle{X: float32((29 * 2) * i), Y: 1, Width: 29 * 2, Height: 36 * 2}, rl.Vector2{X: 0, Y: 0}, 0, rl.White)
		}

		rl.EndDrawing()

	}

	rl.CloseWindow()
}
