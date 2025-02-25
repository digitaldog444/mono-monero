package settings

import (
	"encoding/json"
	"io"
	"os"
)

type Settings struct {
	DSN       string `json:"dsn"`
	Port      int    `json:"port"`
	JWTSecret string `json:"jwt_secret"`
	BankAddress string `json:"bank_address"`
	WithdrawalFee float32 `json:"withdrawal_fee"`
}

func InitializeSettings(file string) (Settings, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return Settings{}, err
	}

	bytes, _ := io.ReadAll(jsonFile)
	var settings Settings
	err = json.Unmarshal(bytes,&settings)
	if err != nil {
		return Settings{}, err
	}
	return settings, nil
}