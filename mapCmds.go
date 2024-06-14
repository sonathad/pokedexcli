package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Name string
	Url  string
}

type MapResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

const (
	nextAreaUrl string = "https://pokeapi.co/api/v2/location-area/"
)

func MapCmd() {
	fmt.Println("Map command used")
	res, err := http.Get(nextAreaUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var mapResponse MapResponse
	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// fmt.Printf("Map Response: %+v\n", mapResponse)
	fmt.Printf("Next is %v\n", string(mapResponse.Next))

}

func MapBCmd() {
	fmt.Println("Mapb command used")
}
