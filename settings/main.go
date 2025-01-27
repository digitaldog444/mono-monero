package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Settings struct {
	DSN       string `json:"dsn"`
	Port      int    `json:"port"`
	JWTSecret string `json:"jwt_secret"`
}

func InitializeSettings(file string) (Settings, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return Settings{}, err
	}

	bytes, _ := ioutil.ReadAll(jsonFile)
	var settings Settings
	err = json.Unmarshal(bytes,&settings)
	if err != nil {
		return Settings{}, err
	}
	return settings, nil
}