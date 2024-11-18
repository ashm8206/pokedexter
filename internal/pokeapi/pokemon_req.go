package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseURl + endpoint

	data, ok := c.cache.Get(fullUrl)
	if ok {
		// fmt.Println("Cache Hit!!!!")
		pokemonResp := Pokemon{}

		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return Pokemon{}, nil
	}
	// fmt.Println("Catch Miss!!!")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, nil
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemonResp := Pokemon{}

	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullUrl, data)
	return pokemonResp, nil

}
