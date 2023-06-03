package config

import (
	"encoding/json"
	"os"
	"log"
)

// Get récupère le contenu du fichier config.json et le stocke dans la variable
// General du package config.
func Get(fileName string) {

	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error when opening config file: ", err)
	}

	err = json.Unmarshal(content, &General)
	if err != nil {
		log.Fatal("Error when reading config file: ", err)
	}
}
