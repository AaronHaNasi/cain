package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	yaml "go.yaml.in/yaml/v4"
)

func getCharacterData(filename string) ([]byte, error) {
	characterData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return characterData, nil
}

func main() {
	var characterData []byte // Holds raw character data before unmarshalling into struct
	var err error            // holds error
	var character Character  // struct to hold character data
	var userInput tea.Model

	if len(os.Args) > 1 { // If filename is provided at launch
		filename := os.Args[1]
		characterData, err = getCharacterData(filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		err = yaml.Unmarshal(characterData, &character)
		if err != nil {
			log.Fatalf("cannot unmarshal data: %v", err)
			os.Exit(2)
		}
	} else { // if filename is not provided at launch
		p := tea.NewProgram(initialModel())
		userInput, err = p.Run()
		if err != nil { // if bubbletea program fails
			log.Fatal(err)
			os.Exit(2)
		} else {
			userInputData := userInput.(model)
			characterData, err = getCharacterData(userInputData.textInput.Value())
			if err != nil { // if file is not opened
				if err == os.ErrNotExist { // if file is not opened because it is not found
					for { // loop until user provides a valid filename
						p = tea.NewProgram(initialModel())
						userInput, err = p.Run()
						if err != nil { // if bubbletea program fails
							log.Fatal(err)
							os.Exit(2)
						} else {
							userInputData = userInput.(model)
							characterData, err = getCharacterData(userInputData.textInput.Value())
							if err != nil { // if file is not opened
								if err == os.ErrNotExist { // if file is not opened because it is not found
									continue
								}
								log.Fatal(err)
								os.Exit(2)
							}
							break
						}

					}
				}
			}
			err = yaml.Unmarshal(characterData, &character)
			if err != nil {
				log.Fatalf("cannot unmarshal data: %v", err)
			}
		}
	}
}
