package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type LocationAreaResponse struct {
	Id                int `json:"id"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	fullUrl := baseURL + "/location-area"
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	data, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreasResp := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResp := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreasResp, nil
}

func (c *Client) ShowLocationArea(locationAreaName string) (LocationAreaResponse, error) {
	if locationAreaName == "" {
		return LocationAreaResponse{}, fmt.Errorf("Empty location area parameter")
	}

	fullUrl := baseURL + "/location-area/" + locationAreaName

	data, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreaResp := LocationAreaResponse{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return LocationAreaResponse{}, fmt.Errorf("Location area with name `%s` was not found", locationAreaName)
	} else if resp.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreaResp, nil
}
