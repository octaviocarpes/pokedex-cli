module github.com/octaviocarpes/pokedex-cli/commands

go 1.23.2

replace github.com/octaviocarpes/pokedex-cli/http-client v0.0.0 => ../http-client

replace github.com/octaviocarpes/pokedex-cli/poke-api v0.0.0 => ../poke-api

replace github.com/octaviocarpes/pokedex-cli/utils v0.0.0 => ../utils

replace github.com/octaviocarpes/pokedex-cli/commands/map v0.0.0 => ./map

require github.com/octaviocarpes/pokedex-cli/poke-api v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/http-client v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/utils v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/commands/map v0.0.0
