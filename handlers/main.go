package handlers

import (
	"github.com/digitaldog444/mono-monero/settings"
	"github.com/gin-gonic/gin"
)

var baseSettings settings.Settings

func InitializeSettings(stgs settings.Settings){
	baseSettings = stgs
}

func HomePage(c *gin.Context){
	c.HTML(200,"index.html", nil)
}