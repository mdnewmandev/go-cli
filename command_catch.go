package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params []string) error {
	pokeName := params[0]
	pokeResp, err := cfg.pokeapiClient.CatchPokemon(pokeName)
	if err != nil {
		return err
	}

	attemptCatch := rand.Intn(pokeResp.BaseExperience + 1)

	fmt.Printf("Throwing a Pokeball at %s...", pokeResp.Name)
	fmt.Println()

	if attemptCatch > 70 {
		fmt.Printf("%s escaped!", pokeResp.Name)
		fmt.Println()
		return nil
	}

	fmt.Printf("%s was caught!", pokeResp.Name)
	fmt.Println()
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.caughtPokemon[pokeResp.Name] = pokeResp
	return nil
}