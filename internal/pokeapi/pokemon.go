package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonResp struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetPokemon(pokemonName string) (PokemonResp, error) {
	if pokemonName == "" {
		return PokemonResp{}, fmt.Errorf("Empty pokemon name parameter")
	}

	fullUrl := baseURL + "/pokemon/" + pokemonName

	data, ok := c.cache.Get(fullUrl)
	if ok {
		pokemonResp := PokemonResp{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return PokemonResp{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return PokemonResp{}, fmt.Errorf("Pokemon with name `%s` was not found", pokemonName)
	} else if resp.StatusCode > 399 {
		return PokemonResp{}, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	pokemonResp := PokemonResp{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	c.cache.Add(fullUrl, data)

	return pokemonResp, nil
}
