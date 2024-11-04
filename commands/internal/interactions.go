package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Store paginations
type config struct {
	prevUrl string
	currentUrl string
	nextUrl string
}

// json to go struct
type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var currConfig = &config {
	nextUrl: "https://pokeapi.co/api/v2/location-area",
}

func GetLocations() (error) {

	// Create get request
	req, err := http.NewRequest("GET", currConfig.nextUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Create http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}

	
	//Decode json
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var locations locationArea
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return fmt.Errorf("no json data")
	}

	// Get locations of all 20 locations in the paginated result
	for _, v := range locations.Results {
		fmt.Println(v.Name)
	}

	// update currConfig url results
	currConfig.currentUrl = currConfig.nextUrl
	if locations.Previous != nil {
		currConfig.prevUrl = locations.Previous.(string)
	}
	currConfig.nextUrl = locations.Next



	return nil
}

func GetLocationsB() (error) {

	if currConfig.prevUrl == "" {
		return fmt.Errorf("previous page not found")
	}

	// Create get request
	req, err := http.NewRequest("GET", currConfig.prevUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Create http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}

	
	//Decode json
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var locations locationArea
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return fmt.Errorf("no json data")
	}

	// Get locations of all 20 locations in the paginated result
	for _, v := range locations.Results {
		fmt.Println(v.Name)
	}

	// update currConfig url results
	if locations.Previous != nil {
		currConfig.prevUrl = locations.Previous.(string)
		currConfig.nextUrl = currConfig.currentUrl
	} 

	currConfig.currentUrl = currConfig.prevUrl

	return nil
}