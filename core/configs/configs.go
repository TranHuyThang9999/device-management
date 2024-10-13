package configs

import (
	"encoding/json"
	"os"
)

type Configs struct {
	DataSource string `json:"data_source"`
	Port       string `json:"port"`
	Secret     string `json:"secret"`
	Expiration string `json:"expiration"`
}

var config *Configs

func Get() *Configs {
	return config
}
func LoadConfig(path string) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	byteValue, err := os.ReadFile(configFile.Name())
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}
}
