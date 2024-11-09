module github.com/octaviocarpes/pokedex-cli/commands

go 1.23.2

replace github.com/octaviocarpes/pokedex-cli/http-client v0.0.0 => ../http-client

replace github.com/octaviocarpes/pokedex-cli/poke-api v0.0.0 => ../poke-api

replace github.com/octaviocarpes/pokedex-cli/pokemon v0.0.0 => ../pokemon

replace github.com/octaviocarpes/pokedex-cli/pokedex v0.0.0 => ../pokedex

replace github.com/octaviocarpes/pokedex-cli/poke-api/cache v0.0.0 => ../poke-api/cache

replace github.com/octaviocarpes/pokedex-cli/utils v0.0.0 => ../utils

replace github.com/octaviocarpes/pokedex-cli/commands/map v0.0.0 => ./map

replace github.com/octaviocarpes/pokedex-cli/commands/explore v0.0.0 => ./explore

replace github.com/octaviocarpes/pokedex-cli/commands/catch v0.0.0 => ./catch

replace github.com/octaviocarpes/pokedex-cli/commands/inspect v0.0.0 => ./inspect

replace github.com/octaviocarpes/pokedex-cli/commands/pokedex v0.0.0 => ./pokedex

require github.com/octaviocarpes/pokedex-cli/poke-api v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/pokemon v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/pokedex v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/poke-api/cache v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/http-client v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/utils v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/commands/map v0.0.0

require github.com/octaviocarpes/pokedex-cli/commands/explore v0.0.0

require github.com/octaviocarpes/pokedex-cli/commands/catch v0.0.0

require github.com/octaviocarpes/pokedex-cli/commands/inspect v0.0.0

require github.com/octaviocarpes/pokedex-cli/commands/pokedex v0.0.0
