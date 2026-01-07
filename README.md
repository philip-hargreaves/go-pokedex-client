# Go Pokédex Client

A command-line Pokédex client built with Go. This tool interacts with the [PokeAPI](https://pokeapi.co/) to allow users to explore different regions, encounter various Pokémon, and maintain their own digital Pokédex.

## Features

- **Interactive REPL**: A user-friendly command-line interface.
- **Location Exploration**: Navigate through different map areas.
- **Pokémon Encounters**: See which Pokémon inhabit specific areas.
- **Catching System**: Attempt to catch Pokémon and add them to your Pokédex.
- **Pokédex Management**: Inspect caught Pokémon and view your entire collection.
- **Efficient Caching**: Internal caching mechanism to reduce API calls and improve performance.

## Installation

Ensure you have [Go](https://go.dev/doc/install) installed on your system.

1. Clone the repository:
   ```bash
   git clone https://github.com/philip-hargreaves/go-pokedex-client.git
   cd go-pokedex-client
   ```

2. Build the project:
   ```bash
   go build -o pokedex
   ```

## Usage

Start the REPL by running the compiled binary:

```bash
./pokedex
```

### Available Commands

Once inside the Pokédex REPL, you can use the following commands:

- `help`: Displays a help message with all available commands.
- `map`: Displays the names of 20 location areas in the Pokémon world. Each subsequent call displays the next 20 areas.
- `mapb`: Displays the previous 20 location areas.
- `explore <location_area>`: Lists all the Pokémon that can be found in a given area.
- `catch <pokemon_name>`: Attempts to catch the specified Pokémon. Catching success is based on the Pokémon's base experience.
- `inspect <pokemon_name>`: Displays details (height, weight, stats, types) of a Pokémon, provided you have caught it.
- `pokedex`: Lists all the Pokémon you have caught so far.
- `exit`: Exits the Pokédex.

## Project Structure

- `main.go`: Entry point of the application.
- `repl.go`: Logic for the REPL loop and command registration.
- `command_*.go`: Implementation of individual CLI commands.
- `internal/pokeapi/`: Client for interacting with the PokeAPI, including type definitions.
- `internal/pokecache/`: A thread-safe, time-based caching system for API responses.

## Architecture

The application is built around a Read-Eval-Print Loop (REPL) that maintains a persistent `config` state across the session.

- **State Management**: A central `config` struct tracks pagination for map commands and stores caught Pokémon in-memory.
- **API Client**: The `pokeapi` package handles all network communication with a custom `Client` that wraps a standard `http.Client`.
- **Caching Layer**: To minimise network requests, the `pokecache` package provides a thread-safe map protected by a `sync.Mutex`. It uses a background goroutine to periodically "reap" (delete) stale entries based on a set TTL (Time-To-Live).
- **Extensibility**: Commands are implemented as callbacks registered in a map, making it easy to add new functionality to the REPL.