package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (RespShallowExplore, error) {
	url := baseURL + "/location-area/" + locationName

	// Check cache first
	if cachedData, ok := c.cache.Get(url); ok {
		fmt.Println("Cache hit!", url)
		exploreResponse := RespShallowExplore{}
		err := json.Unmarshal(cachedData, &exploreResponse)
		if err != nil {
			return RespShallowExplore{}, err
		}
		return exploreResponse, nil
	}
	fmt.Println("Cache miss!", url)

	// If not in cache, make the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowExplore{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowExplore{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowExplore{}, err
	}

	// Cache the raw JSON data
	c.cache.Add(url, dat)

	exploreResponse := RespShallowExplore{}
	err = json.Unmarshal(dat, &exploreResponse)
	if err != nil {
		return RespShallowExplore{}, err
	}

	return exploreResponse, nil
}