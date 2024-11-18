package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	// "golang.org/x/text/cases"
	// "golang.org/x/text/language"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	_, ok := cfg.caughtPokemon[pokemonName]
	if ok {
		return fmt.Errorf("pokemon already caught")
	}

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		log.Fatal(err)
	}

	const threshold = 50

	// rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time
	// strconv.Atoi(s)
	// fmt.Println(reflect.TypeOf(resp.BaseExperience))
	randNumber := rand.Intn(21) + 40

	// exp := resp.BaseExperience
	// fmt.Println(exp)

	// randNumber := rand.Intn(exp)
	if randNumber < threshold {
		return fmt.Errorf("failed to catch the pokemon %s", pokemonName)
	}
	cfg.caughtPokemon[pokemonName] = resp

	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("*********************************")
	fmt.Printf("Pokemon Strength:  %v \n", resp.BaseExperience)
	fmt.Printf("Strength Needed:  %v \n", threshold)
	fmt.Printf("Your Strength:  %v \n", randNumber)

	return nil

}
