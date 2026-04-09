package main 

import(
	"fmt"
	"os"
	"strings"
	"encoding/json"
)


type Pokemon struct {
	Name					string
	Type					string
	Health				int
	Defense				int
	SpAttack			int 
	SpDefense 		int 
	Speed					int
}


func main() {

	pokemon := load()

	var userChoice, userDel string


	for {

		fmt.Println("What would you like to do?\nView?\nAdd/Update?\nRemove?\nSave?\nExit?")

		fmt.Scanln(&userChoice)

		switch strings.ToLower(userChoice) {
		case "view":
			view(pokemon)
		case "add", "update":
			add(pokemon)
		case "remove", "delete":
			fmt.Println("What would you like to remove?\n")
			fmt.Scanln(&userDel)
			delete(pokemon, strings.ToLower(userDel))
		case "save":
			save(pokemon)
		case "exit", "close":
			os.Exit(0)
		}
	}
}


func view(pokemon map[string]Pokemon) {
	for name, p := range pokemon{
		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Type: %s\n", p.Type)
		fmt.Printf("Health: %d\n", p.Health)
		fmt.Printf("Defense: %d\n", p.Defense)
		fmt.Printf("SpAttack: %d\n", p.SpAttack)
		fmt.Printf("SpDefense: %d\n", p.SpDefense)
		fmt.Printf("Speed: %d\n", p.Speed)
	}
}



func add(pokemon map[string]Pokemon) {
	var inputHealth, inputDefense, inputSpattack, inputSpdefense, inputSpeed int
	var inputName, inputType string
	fmt.Println("Enter/Change in this order: Name, Type, Health, Defense, SpAttack, SpDefense, Speed:\n")
	fmt.Scanln(&inputName, &inputType, &inputHealth, &inputDefense, &inputSpattack, &inputSpdefense, &inputSpeed)
	pokemon[inputName] = Pokemon{Name: inputName, Type: inputType, Health: inputHealth, Defense: inputDefense, SpAttack: inputSpattack, SpDefense: inputSpdefense, Speed: inputSpeed}
}


func save(pokemon map[string]Pokemon) {
	data, err := json.Marshal(pokemon)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	err = os.WriteFile("Pokedex.json", data, 0644)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	fmt.Println("Saved Successfully")
}



func load() map[string]Pokemon{
	data, err := os.ReadFile("Pokedex.json")
	if err != nil {
		fmt.Println("Error Reading File:", err)
		return make(map[string]Pokemon)
	}

	var loadedDex map[string]Pokemon
	err = json.Unmarshal(data, &loadedDex)
	if err != nil {
		fmt.Println("Error Unmarshaling data:", err)
		return make(map[string]Pokemon)
	}
	return loadedDex
}

