package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	// path
	endpoint := "/pokemon/" + pokemonName
	// full endpoint
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		pokemon := Pokemon{}
		// convert json to structs
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}
	fmt.Println("cache miss!")

	// wrap to background
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// make api call
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	// close resposne body after
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// convert body to bytes
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	// convert json to structs
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
