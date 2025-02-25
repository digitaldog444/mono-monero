package main

import (
	"fmt"

	"github.com/digitaldog444/mono-monero/database"
	"github.com/digitaldog444/mono-monero/handlers"
	"github.com/digitaldog444/mono-monero/routes"
	"github.com/digitaldog444/mono-monero/settings"
)

func main() {
	baseSettings, err := settings.InitializeSettings("./settings.json")
	if err != nil {
		fmt.Println("Error occured while initializing settings!")
		fmt.Println(err)
	}
	db := database.InitializeDatabase(baseSettings.DSN)
	handlers.InitializeDB(db)

	handlers.InitializeSettings(baseSettings)
	routes.InitializeRouter(baseSettings.Port)
}