package client

import (
	"net/http"
	"time"

	"github.com/martinpare1208/pokedexcli/internal/cache"
	"github.com/martinpare1208/pokedexcli/internal/pokedex"
)

type Client struct {
	Cache pokecache.Cache
	HttpClient http.Client
	ClientPokedex map[string]pokedex.Pokemon
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(cacheInterval),
		HttpClient: http.Client{
			Timeout: timeout,
		},
		ClientPokedex: make(map[string]pokedex.Pokemon),
	}
}