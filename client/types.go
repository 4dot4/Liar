package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type screens int32

type typeCard struct {
	CardValue int32
	CardTeam  int32
}
type Card struct {
	TypeCard  typeCard
	RecSource rl.Rectangle
}

const (
	Width  = 800
	Heigth = 800
)

// cardteams
const (
	Clubs = iota
	Diamonds
	Hearts
	Spades
	back
)

// card valuoes
const (
	As    = 1
	J     = 11
	Q     = 12
	K     = 13
	scale = 2
)

// screens
const (
	Start screens = iota
	Client
	Server
)

type player struct {
	Name string
	Id   int
}
type game struct {
	RayLogo     rayLogo
	GameScreen  screens
	Spritesheet rl.Texture2D
	Map         map[typeCard]rl.Rectangle
	Hand        []typeCard
}
type rayLogo struct {
	logoPositionX      int32
	logoPositionY      int32
	framesCounter      int32
	lettersCount       int32
	topSideRecWidth    int32
	leftSideRecHeight  int32
	bottomSideRecWidth int32
	rightSideRecHeight int32
	state              int32
	alpha              float32
}
