package main

import "fmt"

func commandExplore(cfg *config, params []string) error {
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(params[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s", exploreResp.Name)
	fmt.Println()
	fmt.Println("Found Pokemon:")

	for _, encounters := range exploreResp.PokemonEncounters {
		fmt.Println("- ", encounters.Pokemon.Name)
	}

	return nil
}