package main

import (
	"errors"
	"fmt"
)

func commandMapNext(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = res.Next
	cfg.prevLocationsURL = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapPrev(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	res, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = res.Next
	cfg.prevLocationsURL = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}
