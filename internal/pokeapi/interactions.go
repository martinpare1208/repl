package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/martinpare1208/pokedexcli/internal/config"
)

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


func GetLocations(cfg *config.Cfg) (error) {

	fullURL := cfg.NextUrl

	if cfg.NextUrl == "" {
		fullURL = baseURL + "/location-area"
	}


	

	// Create get request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Send request
	resp, err := cfg.PokeClient.HttpClient.Do(req)
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

	// update cfg url results
	cfg.CurrentUrl = cfg.NextUrl
	if locations.Previous != nil {
		cfg.PrevUrl = locations.Previous.(string)
	}
	cfg.NextUrl = locations.Next
	


	return nil
}

func GetLocationsB(cfg *config.Cfg) (error) {

	if cfg.PrevUrl == "" {
		return fmt.Errorf("previous page not found")
	}

	// Create get request
	req, err := http.NewRequest("GET", cfg.PrevUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Send the request
	resp, err := cfg.PokeClient.HttpClient.Do(req)
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

	// update cfg url results
	if locations.Previous != nil {
		cfg.PrevUrl = locations.Previous.(string)
		cfg.NextUrl = cfg.CurrentUrl
	} 

	cfg.CurrentUrl = cfg.PrevUrl

	return nil
}