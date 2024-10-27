module github.com/octaviocarpes/pokedex-cli/poke-api

go 1.23.2

replace github.com/octaviocarpes/pokedex-cli/http-client v0.0.0 => ../http-client

require (
	github.com/octaviocarpes/pokedex-cli/http-client v0.0.0
)