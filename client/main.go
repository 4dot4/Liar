package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var Game game
	const (
		Width  = 800
		Heigth = 800
	)
	Game.GameScreen = Start
	Game.RayLogo = rayLogo{logoPositionX: int32(Width/2 - 128), logoPositionY: int32(Heigth/2 - 128), framesCounter: 0, lettersCount: 0, topSideRecWidth: 16, leftSideRecHeight: 16, bottomSideRecWidth: 16, rightSideRecHeight: 16, alpha: 1}

	//channel := make(chan game, 5)

	Game.Map =  initSprites()

	rl.InitWindow(Width, Heigth, "The Liar")
	rl.SetTargetFPS(60)
	go client()
	Game.Spritesheet = rl.LoadTexture("../assets/spritesheet.png")

	for !rl.WindowShouldClose() {

		switch Game.GameScreen {
		case Client:

		case Server:

		}

		update(&Game)
		drawing(&Game, &Game.Hand)

	}
	rl.UnloadTexture(Game.Spritesheet)
	rl.CloseWindow()
}
