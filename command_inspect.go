package main

import "fmt"

func commandInspect(cfg *config, params []string) error {
	pokeName := params[0]
	pokeResp, exists := cfg.caughtPokemon[pokeName]
	if !exists {
		return fmt.Errorf("you have not caught a Pokemon named %s", pokeName)
	}

	fmt.Printf("Name: %s", pokeResp.Name)
	fmt.Println()
	fmt.Printf("Height: %d", pokeResp.Height)
	fmt.Println()
	fmt.Printf("Weight: %d", pokeResp.Weight)
	fmt.Println()
	fmt.Printf("Stats:")
	fmt.Println()
	for _, stat := range pokeResp.Stats {
		fmt.Printf("- %s: %d", stat.Stat.Name, stat.BaseStat)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Types:")
	for _, t := range pokeResp.Types {
		fmt.Printf("- %s", t.Type.Name)
		fmt.Println()
	}

	return nil
}