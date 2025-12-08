package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	yaml "go.yaml.in/yaml/v4"
)

func getCharacterData(filename string) []byte {
	characterData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return characterData
}

func main() {
	var character Character
	if len(os.Args) > 1 {
		filename := os.Args[1]
		characterData := getCharacterData(filename)
		err := yaml.Unmarshal(characterData, &character)
		if err != nil {
			log.Fatalf("cannot unmarshal data: %v", err)
		}
	} else {
		p := tea.NewProgram(initialOpenFileModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
			os.Exit(2)
		}
	}
}
