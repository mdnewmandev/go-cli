package main

func commandPokedex(cfg *config, params []string) error {
	if len(cfg.caughtPokemon) == 0 {
		println("You have not caught any Pokemon yet.")
		return nil
	}

	println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		println("- " + name)
	}

	return nil
}