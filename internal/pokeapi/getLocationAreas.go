package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl *string) (ResponseLocationAreas, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	locationAreas := ResponseLocationAreas{}
	if data, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return ResponseLocationAreas{}, err
		}

		return locationAreas, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocationAreas{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	c.cache.Add(url, data)
	return locationAreas, nil
}
