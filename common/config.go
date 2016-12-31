package common

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Symbols []string
}

var configuration Configuration

func init() {
	configuration = Configuration{}
	loadConfig(&configuration)
}

func GetConfig() Configuration {
	return configuration
}
func loadConfig(configuration *Configuration) {
	file, err := os.Open("common/config.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}
