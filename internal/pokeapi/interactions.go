package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/martinpare1208/pokedexcli/internal/config"
	"github.com/martinpare1208/pokedexcli/internal/pokedex"
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


type locationAreaInformation struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func GetPokemonInArea(cfg *config.Cfg, location string) (error) {
	fullURL := baseURL + "/location-area/" + location

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

	var locationAreaInformation locationAreaInformation
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locationAreaInformation)
	if err != nil {
		return fmt.Errorf("error: json could not be decoded")
	}

	for _, v := range locationAreaInformation.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}

	return nil
}


func CatchPokemon(cfg *config.Cfg, pokemon string) (error) {

	fullURL := baseURL + "/pokemon/" + pokemon

	//Create get request
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return err
	}

	//Send get request
	resp, err := cfg.PokeClient.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error: could not complete request: %w", err)
	}

	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: unexpected status code %d", resp.StatusCode)
	}

	var pokemonInfo pokedex.Pokemon
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&pokemonInfo)
	if err != nil {
		return fmt.Errorf("could not decode json")
	}

	fmt.Printf("You threw a pokeball at %s", pokemon)
	fmt.Println()
	catch := GenerateSuccessRate(pokemonInfo.BaseExperience)
	if catch {
		fmt.Printf("You caught %s!", pokemon)
		fmt.Println()
		_, exists := cfg.PokeClient.ClientPokedex[pokemon]
		if exists {
			fmt.Println("You already caught this pokemon good job!")
			return nil
		}

		// Add to pokedex
		fmt.Println("Adding to your pokedex!")
		cfg.PokeClient.ClientPokedex[pokemon] = pokemonInfo

	} else {
		fmt.Println("Catch failed! Try again!")
	}



	return nil
}


func PrintCurrentPokedex(cfg *config.Cfg) (error) {
	for _, v := range cfg.PokeClient.ClientPokedex {
		fmt.Println(v.Name)
	}
	return nil
}