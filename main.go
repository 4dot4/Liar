package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Card struct {
	Number   cardValue
	TypeCard typeCard
	RecPos   rl.Vector2
}

func main() {
	var cards [4][13]Card

	for y := 0; y < 4; y++ {
		for x := 0; x < 13; x++ {
			switch y {
			case int(Clubs):
				cards[y][x].TypeCard = Clubs
			case int(Diamonds):
				cards[y][x].TypeCard = Diamonds
			case int(Hearts):
				cards[y][x].TypeCard = Hearts
			case int(Spades):
				cards[y][x].TypeCard = Spades

			}
			cards[y][x].Number = cardValue(x)
		}
	}
	fmt.Println(cards)
	rl.InitWindow(800, 800, "Liar")
	var Spritesheet rl.Texture2D = rl.LoadTexture("./assets/spritesheet.png")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		 rl.DrawTexture(Spritesheet,0,0,rl.White)
		rl.DrawText("HELLO WORLD", 400, 400, 20, rl.Black)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
