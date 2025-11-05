package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (RespShallowPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// Check cache first
	if cachedData, ok := c.cache.Get(url); ok {
		fmt.Println("Cache hit!", url)
		catchResponse := RespShallowPokemon{}
		err := json.Unmarshal(cachedData, &catchResponse)
		if err != nil {
			return RespShallowPokemon{}, err
		}
		return catchResponse, nil
	}
	fmt.Println("Cache miss!", url)

	// If not in cache, make the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	// Cache the raw JSON data
	c.cache.Add(url, dat)

	catchResponse := RespShallowPokemon{}
	err = json.Unmarshal(dat, &catchResponse)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	return catchResponse, nil
}