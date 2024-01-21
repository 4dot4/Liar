package main

import (
	"fmt"
	"log"
	"net/http"
)

func cardsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cards" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!\n")
}

func server(ch <-chan []typeCard) {
	http.HandleFunc("/cards", cardsHandler) // Update this line of code
	
	
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
