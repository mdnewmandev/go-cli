package pokeapi

import (
	"net/http"
	"time"

	"github.com/mdnewmandev/go-cli/internal/pokecache"
)

// Client
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// New Client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Minute),
	}
}