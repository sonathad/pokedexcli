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
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

var mapResponse MapResponse

func MapCmd(areaUrl string) {
	if areaUrl == "next" {
		if mapResponse.Next != nil {
			areaUrl = *mapResponse.Next
		} else {
			areaUrl = "https://pokeapi.co/api/v2/location-area/"
		}
	} else if areaUrl == "prev" {
		areaUrl = *mapResponse.Previous
	} else {
		areaUrl = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(areaUrl)
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

	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	for _, location := range mapResponse.Results {
		fmt.Println(location.Name)
	}
}
