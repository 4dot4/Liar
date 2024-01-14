package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawing(Spritesheet *rl.Texture2D, cards *[4][13]Card, GameScreen *screens) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	for i := 0; i < 4; i++ {
		for _, v := range cards[i] {

			rl.DrawTexturePro(*Spritesheet, v.RecSource, rl.Rectangle{X: float32(29 * i * 2), Y: float32(36 * i * 2), Width: 29 * 2, Height: 36 * 2}, rl.Vector2{0, 0}, 0, rl.White)
		}

		//rl.DrawTexturePro(Spritesheet, cards[x][i].RecSource, rl.Rectangle{X: float32((29 * 2) * i), Y: 1, Width: 29 * 2, Height: 36 * 2}, rl.Vector2{X: 0, Y: 0}, 0, rl.White)
	}

	rl.EndDrawing()
}
