package main

import (
	

	rl "github.com/gen2brain/raylib-go/raylib"
)
type typeCard int32
type cardValue int32
const (
	Clubs typeCard = iota
	Diamonds 
	Hearts 
	Spades
	As cardValue = 1
	J = 11
	Q = 12 
	K = 13
)
type Card struct {
	Number cardValue
	TypeCard typeCard
	RecPos rl.Vector2	
}

func main() {
	var cards [4][13]Card
	
	for y := 0 ; y < 4; y++{
		for x := 0 ; x < 13; x++{
			switch y { 
				case int(Clubs) :
					cards[y][x].TypeCard = Clubs
				case int(Diamonds) :
					cards[y][x].TypeCard = Diamonds
				case int(Hearts) :
					cards[y][x].TypeCard = Hearts
				case int(Spades) :
					cards[y][x].TypeCard = Spades

			}
			cards[y][x].Number = cardValue(x)
		}
	}

	
	rl.InitWindow(800,800,"Liar")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText("HELLO WORLD",400,400,20,rl.Black)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
