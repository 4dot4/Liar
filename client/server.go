package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func cardsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cards" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	hand := []typeCard{{As, Hearts}, {2, Spades}}
	jsonHand, _ := json.Marshal(hand)

	fmt.Fprintf(w, string(jsonHand))
}

func server() {

	http.HandleFunc("/cards", cardsHandler) // Update this line of code

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
