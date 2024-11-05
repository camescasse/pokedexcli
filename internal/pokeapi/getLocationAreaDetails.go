package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaDetails(areaName string) (ResponseLocationAreaDetails, error) {
	if areaName == "" {
		return ResponseLocationAreaDetails{}, errors.New("location area name required")
	}
	url := baseUrl + "/location-area/" + areaName

	locationAreaDetails := ResponseLocationAreaDetails{}
	if data, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(data, &locationAreaDetails)
		if err != nil {
			return ResponseLocationAreaDetails{}, err
		}

		return locationAreaDetails, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocationAreaDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocationAreaDetails{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocationAreaDetails{}, err
	}

	err = json.Unmarshal(data, &locationAreaDetails)
	if err != nil {
		return ResponseLocationAreaDetails{}, errors.New("invalid location area")
	}

	c.cache.Add(url, data)
	return locationAreaDetails, nil
}
