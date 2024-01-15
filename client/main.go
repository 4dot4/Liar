package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var game game
	const (
		Width  = 800
		Heigth = 800
	)
	game.GameScreen = Start
	game.RayLogo = rayLogo{logoPositionX: int32(Width/2 - 128), logoPositionY: int32(Heigth/2 - 128), framesCounter: 0, lettersCount: 0, topSideRecWidth: 16, leftSideRecHeight: 16, bottomSideRecWidth: 16, rightSideRecHeight: 16, alpha: 1}

	channel := make(chan [4][13]Card, 5)
	go client(channel)
	initSprites(&game.Cards)

	rl.InitWindow(800, 800, "The Liar")
	rl.SetTargetFPS(60)

	game.Spritesheet = rl.LoadTexture("../assets/spritesheet.png")

	for !rl.WindowShouldClose() {

		update(&game)

		drawing(&game)

	}
	rl.UnloadTexture(game.Spritesheet)
	rl.CloseWindow()
}
