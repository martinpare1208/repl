package client

import (
	"net/http"
	"time"
	"github.com/martinpare1208/pokedexcli/internal/cache"
)

type Client struct {
	Cache pokecache.Cache
	HttpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(cacheInterval),
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}