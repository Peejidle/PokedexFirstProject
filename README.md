# Pokedex CLI

A command-line Pokedex built in Go. Add, view, and remove Pokemon entries that persist between sessions via JSON storage.

## Features

- Add or update Pokemon with full stat tracking
- View all saved Pokemon
- Remove entries by name
- Automatic save/load to `Pokedex.json`

## Usage
```bash
go run pokedex.go
```

### Commands

| Command | Description |
|---------|-------------|
| `view` | Display all saved Pokemon |
| `add` | Add or update a Pokemon |
| `remove` | Remove a Pokemon by name |
| `save` | Save your Pokedex to disk |
| `exit` | Exit the program |

When adding a Pokemon, enter values in this order:
Name Type Health Defense SpAttack SpDefense Speed

## Example
Add
Enter/Change in this order: Name, Type, Health, Defense, SpAttack, SpDefense, Speed:
Bulbasaur Grass 45 49 65 65 45


