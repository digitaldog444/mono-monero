package main

import (
	"fmt"

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
	handlers.InitializeSettings(baseSettings)
	routes.InitializeRouter(baseSettings.Port)
}