package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"net/http"
	"time"
)

func client() {
	fmt.Println("ADENTREI DESGRACA")
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", "http://localhost:8080/cards", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	var hand []typeCard
	if er := json.Unmarshal(body, &hand); er != nil {
		log.Fatal(er)
	}
	fmt.Println(hand)

	fmt.Printf("Response status : %s \n", resp.Status)
}