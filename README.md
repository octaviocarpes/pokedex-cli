# Pokedex CLI

A command-line interface (CLI) Pokedex application built with Go. This CLI app allows you to explore locations in the Pokémon world, attempt to catch Pokémon, and view your own Pokedex collection.

## Features

- **Help**: Get instructions on how to use the Pokedex CLI.
- **Map**: View available locations in the Pokémon world.
- **Explore**: Discover Pokémon in a specific location.
- **Catch**: Try to catch a Pokémon you encounter.
- **Pokedex**: View your collection of caught Pokémon.

## Prerequisites

- **Go**: You need to have Go installed on your system.

## Installing Go

### 1. Download Go

Visit the [official Go website](https://go.dev/dl/) and download the installer for your OS.

### 2. Install Go

Follow the installation instructions for your operating system.

### 3. Verify Installation

Open a terminal and run the following command:

```
go version
```

You should see output like `go version go1.xx.x`.

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/octaviocarpes/pokedex-cli.git
cd pokedex-cli
```

### Build the Project

```bash
go build -o pokedex
```

### Start the Pokedex CLI

```bash
./pokedex
```

## Available Commands

Here is a list of all available commands:

### `help`

Displays information on how to use the Pokedex CLI.

**Example**:
```bash
Pokedex CLI > help
```

### `exit`

Exits the Pokedex CLI program.

**Example**:
```bash
Pokedex CLI > exit
```

### `map`

Fetches the locations of the Pokémon world (in forward order).

**Example**:
```bash
Pokedex CLI > map
```

### `mapb`

Fetches the locations of the Pokémon world (in reverse order).

**Example**:
```bash
Pokedex CLI > mapb
```

### `explore <location>`

Explore a specific location in the Pokémon world to see a list of Pokémon you can encounter.

**Example**:
```bash
Pokedex CLI > explore mt-coronet-2f
```

### `catch <pokemon>`

Attempt to catch a specific Pokémon.

**Example**:
```bash
Pokedex CLI > catch pikachu
```

### `pokedex`

Displays the Pokémon you have caught.

**Example**:
```bash
Pokedex CLI > pokedex
```

## Contributing

Feel free to contribute to this project! Fork the repository, create a new branch, and submit a pull request with your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Pokémon is a registered trademark of Nintendo, Game Freak, and Creatures. This project is a fan-made application and is not affiliated with any of the official Pokémon companies.

