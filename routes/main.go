package routes

import (
	"strconv"

	"github.com/digitaldog444/mono-monero/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(port int) *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	r.Static("./static","static")
	r.GET("/",handlers.HomePage)
	r.Run(":" + strconv.Itoa(port))
	return r
}