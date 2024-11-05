package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (ResponsePokemon, error) {
	if name == "" {
		return ResponsePokemon{}, errors.New("pokemon name required")
	}
	url := baseUrl + "/pokemon/" + name

	pokemon := ResponsePokemon{}
	if data, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return ResponsePokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponsePokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponsePokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponsePokemon{}, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return ResponsePokemon{}, errors.New("invalid pokemon")
	}

	c.cache.Add(url, data)
	return pokemon, nil
}
