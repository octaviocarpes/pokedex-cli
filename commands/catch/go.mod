module github.com/octaviocarpes/pokedex-cli/commands/catch

go 1.23.2

replace github.com/octaviocarpes/pokedex-cli/http-client v0.0.0 => ../../http-client

replace github.com/octaviocarpes/pokedex-cli/poke-api v0.0.0 => ../../poke-api

replace github.com/octaviocarpes/pokedex-cli/pokedex v0.0.0 => ../../pokedex

replace github.com/octaviocarpes/pokedex-cli/pokemon v0.0.0 => ../../pokemon

replace github.com/octaviocarpes/pokedex-cli/poke-api/cache v0.0.0 => ../../poke-api/cache

replace github.com/octaviocarpes/pokedex-cli/utils v0.0.0 => ../../utils

require github.com/octaviocarpes/pokedex-cli/http-client v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/pokedex v0.0.0

require github.com/octaviocarpes/pokedex-cli/pokemon v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/poke-api v0.0.0

require github.com/octaviocarpes/pokedex-cli/poke-api/cache v0.0.0 // indirect

require github.com/octaviocarpes/pokedex-cli/utils v0.0.0 // indirect