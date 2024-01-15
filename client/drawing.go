package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Drawintro(r *rayLogo) {
	if r.state == 0 {
		if (r.framesCounter/15)%2 == 1 {
			rl.DrawRectangle(r.logoPositionX, r.logoPositionY, 16, 16, rl.Black)
		}
	} else if r.state == 1 {
		rl.DrawRectangle(r.logoPositionX, r.logoPositionY, r.topSideRecWidth, 16, rl.Black)
		rl.DrawRectangle(r.logoPositionX, r.logoPositionY, 16, r.leftSideRecHeight, rl.Black)
	} else if r.state == 2 {
		rl.DrawRectangle(r.logoPositionX, r.logoPositionY, r.topSideRecWidth, 16, rl.Black)
		rl.DrawRectangle(r.logoPositionX, r.logoPositionY, 16, r.leftSideRecHeight, rl.Black)

		rl.DrawRectangle(r.logoPositionX+240, r.logoPositionY, 16, r.rightSideRecHeight, rl.Black)
		rl.DrawRectangle(r.logoPositionX, r.logoPositionY+240, r.bottomSideRecWidth, 16, rl.Black)
	} else if r.state == 3 {
		rl.DrawRectangle(r.logoPositionX, r.logoPositionY, r.topSideRecWidth, 16, rl.Fade(rl.Black, r.alpha))
		rl.DrawRectangle(r.logoPositionX, r.logoPositionY+16, 16, r.leftSideRecHeight-32, rl.Fade(rl.Black, r.alpha))

		rl.DrawRectangle(r.logoPositionX+240, r.logoPositionY+16, 16, r.rightSideRecHeight-32, rl.Fade(rl.Black, r.alpha))
		rl.DrawRectangle(r.logoPositionX, r.logoPositionY+240, r.bottomSideRecWidth, 16, rl.Fade(rl.Black, r.alpha))
		text := "raylib"
		length := int32(len(text))

		rl.DrawRectangle(int32(rl.GetScreenWidth()/2-112), int32(rl.GetScreenHeight()/2-112), 224, 224, rl.Fade(rl.RayWhite, r.alpha))
		if r.lettersCount > length {
			r.lettersCount = length
		}
		rl.DrawText(text[0:r.lettersCount], int32(rl.GetScreenWidth()/2-44), int32(rl.GetScreenHeight()/2+48), 50, rl.Fade(rl.Black, r.alpha))
	} else if r.state == 4 {
		rl.DrawText("[R] REPLAY", 340, 200, 20, rl.Gray)
	}

}
func drawClient(game *game) {

}
func drawing(game *game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	switch game.GameScreen {
	case Start:
		Drawintro(&game.RayLogo)
	case Client:
		drawClient(game)
	case Server:

	}

	rl.EndDrawing()
}
