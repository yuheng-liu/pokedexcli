package pokeapi

import (
	"net/http"
	"time"

	"github.com/yuheng-liu/pokedexcli/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// struct representing local client to make https calls
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		// adds cache function to local client
		cache: pokecache.NewCache(cacheInterval),
		// standard library http that allows configuring of timeout
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
