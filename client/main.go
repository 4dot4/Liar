package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var Game game

	go server()
	go client()
	Game.GameScreen = Start

	Game.Map = initSprites()

	rl.InitWindow(Width, Heigth, "CARDS")
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)

	Game.RayLogo = rayLogo{logoPositionX: int32(rl.GetScreenWidth()/2 - 128), logoPositionY: int32(rl.GetScreenHeight()/2 - 128), framesCounter: 0, lettersCount: 0, topSideRecWidth: 16, leftSideRecHeight: 16, bottomSideRecWidth: 16, rightSideRecHeight: 16, alpha: 1}
	Game.Spritesheet = rl.LoadTexture("../assets/spritesheet.png")

	for !rl.WindowShouldClose() {

		switch Game.GameScreen {
		case Start:
			if rl.IsWindowResized() {
				Game.RayLogo.logoPositionX = int32(rl.GetScreenWidth()/2 - 128)
				Game.RayLogo.logoPositionY = int32(rl.GetScreenHeight()/2 - 128)
			}

		case Client:

		case Server:

		}

		update(&Game)
		drawing(&Game, &Game.Hand)

	}
	rl.UnloadTexture(Game.Spritesheet)
	rl.CloseWindow()
}

//this is a fucking test and i dont know what im doing
