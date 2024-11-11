package config

import "github.com/martinpare1208/pokedexcli/internal/client"

type Cfg struct {
	PokeClient client.Client
	PrevUrl string
	CurrentUrl string
	NextUrl string

}