package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache first
	if cachedData, ok := c.cache.Get(url); ok {
		fmt.Println("Cache hit!", url)
		locationsResponse := RespShallowLocations{}
		err := json.Unmarshal(cachedData, &locationsResponse)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResponse, nil
	}
	fmt.Println("Cache miss!", url)

	// If not in cache, make the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Cache the raw JSON data
	c.cache.Add(url, dat)

	locationsResponse := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResponse)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResponse, nil
}