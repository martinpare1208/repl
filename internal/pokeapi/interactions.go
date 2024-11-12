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
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func GetLocations(cfg *config.Cfg, url string) (error) {

	// Check if url in cache

	data, exists := cfg.PokeClient.Cache.Get(url)
	if exists {
		var locationsCache locationArea
		err := json.Unmarshal(data, &locationsCache)
		if err == nil {
			fmt.Println("key found!")
			for _, v := range locationsCache.Results {
				fmt.Println(v.Name)
			}
			cfg.NextUrl = locationsCache.Next
			cfg.PrevUrl = locationsCache.Previous
			return nil
		}

	}

	

	fullURL := url

	if url == "" {
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

	dat, err := json.Marshal(locations)
	if err != nil {
		return err
	}
	

	// Get locations of all 20 locations in the paginated result
	for _, v := range locations.Results {
		fmt.Println(v.Name)
	}


	//Add to cache
	fmt.Println("Adding link to cache")
	cfg.PokeClient.Cache.Add(fullURL, dat)



	// update cfg url results
	cfg.NextUrl = locations.Next
	cfg.PrevUrl = locations.Previous

	return nil
}

func GetLocationsB(cfg *config.Cfg, url string) (error) {
	if url == "" {
		return fmt.Errorf("error: already on first page")
	}
	GetLocations(cfg, url)
	return nil
}

