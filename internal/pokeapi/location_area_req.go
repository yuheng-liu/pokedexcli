package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client struct method, to make the http call and retrieve data
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	// path
	endpoint := "/location-area"
	// full endpoint
	fullURL := baseURL + endpoint

	// used for when prev/next pages exists, will use those endpoints instead
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache, if entry exists
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		locationAreasResp := LocationAreasResp{}
		// Unmarshalling, converting json objects received into structs
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("cache miss!")

	// Make a new request that moves the operation to background, this does not call the api yet
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// calls the api here and get a response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	// always remember to close the response body
	defer resp.Body.Close()

	// status code 400 and above are not handled internally
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// convert body of response to array of bytes
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	// Unmarshalling, converting json objects received into structs
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// add data to cache
	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	// path
	endpoint := "/location-area/" + locationAreaName
	// full endpoint
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}
	fmt.Println("cache miss!")

	// wrap to background
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// make the api call
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	// close response body at the end
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// convert body to bytes
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	// Unmarshalling, converting json objects received into structs
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	// add to cache
	c.cache.Add(fullURL, data)

	return locationArea, nil
}
