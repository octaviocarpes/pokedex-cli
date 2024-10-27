module github.com/octaviocarpes/pokedex-cli

go 1.23.2

replace github.com/octaviocarpes/pokedex-cli/commands v0.0.0 => ./commands

replace github.com/octaviocarpes/pokedex-cli/poke-api v0.0.0 => ./poke-api

replace github.com/octaviocarpes/pokedex-cli/http-client v0.0.0 => ./http-client

require github.com/octaviocarpes/pokedex-cli/commands v0.0.0

require github.com/octaviocarpes/pokedex-cli/poke-api v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/http-client v0.0.0 // indirect
