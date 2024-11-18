package main

import (
	"fmt"
	"log"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("Too many or too little locations provided")

	}
	locationArea := args[0]
	resp, err := cfg.pokeapiClient.GetLocationAreas(locationArea)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pokemons in the Area")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" -- %s\n", pokemon.Pokemon.Name)
	}
	return nil

}
