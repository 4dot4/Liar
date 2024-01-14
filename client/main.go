package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	channel := make(chan [4][13]Card, 5)
	go client(channel)
	var GameScreen screens = Start
	var cards [4][13]Card

	initSprites(&cards)
	fmt.Println(cards)
	rl.InitWindow(800, 800, "Liar")
	rl.SetTargetFPS(60)

	var Spritesheet rl.Texture2D = rl.LoadTexture("../assets/spritesheet.png")

	for !rl.WindowShouldClose() {

		//update()

		drawing(&Spritesheet, &cards, &GameScreen)

	}

	rl.CloseWindow()
}
