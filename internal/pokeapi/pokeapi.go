package pokeapi

import (
	"net/http"
	"time"

	"github.com/rkapu/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}
