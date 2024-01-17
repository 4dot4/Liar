package main

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