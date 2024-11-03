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

	locationAreas := ResponseLocationAreas{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	return locationAreas, nil
}
