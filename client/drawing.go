package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawClient(game *game) {
	for k, v := range game.Hand {
		rl.DrawTexturePro(game.Spritesheet, game.Map[v], rl.Rectangle{X: float32(k * 29 * scale), Y: 40, Width: 29 * scale, Height: 36 * scale}, rl.Vector2{0, 0}, 0, rl.White)
	}
}
func drawing(game *game, hand *[]typeCard) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.DarkBlue)
	switch game.GameScreen {
	case Start:
		Drawintro(&game.RayLogo)
	case Client:
		drawClient(game)
		rl.DrawText("CARDS", int32(rl.GetScreenWidth()/2)-150, int32(rl.GetScreenHeight()/2), 60, rl.White)
	case Server:

	}

	rl.EndDrawing()
}

// raylib logo intro DO NOT OPEN NEVER
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
