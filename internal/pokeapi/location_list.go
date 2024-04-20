package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationsResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResponse{}, err
	}

	locationsResp := LocationsResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationsResponse{}, err
	}

	return locationsResp, nil
}