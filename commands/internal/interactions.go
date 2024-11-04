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
	currentUrl: "https://pokeapi.co/api/v2/location-area",
}

func GetLocations() (error) {
	// Create get request
	req, err := http.NewRequest("GET", currConfig.currentUrl, nil)
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
	currConfig.prevUrl = currConfig.currentUrl
	currConfig.currentUrl = locations.Next

	return nil
}