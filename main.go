package main

import (
	"time"

	"github.com/mdnewmandev/go-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.RespShallowPokemon),
	}

	startRepl(cfg)
}