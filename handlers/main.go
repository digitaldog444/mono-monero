package handlers

import (
	"github.com/digitaldog444/mono-monero/models"
	"github.com/digitaldog444/mono-monero/settings"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var baseSettings settings.Settings
var db *gorm.DB
func InitializeSettings(stgs settings.Settings){
	baseSettings = stgs
}

func InitializeDB(dbase *gorm.DB) {
	db = dbase
}

func HomePage(c *gin.Context){
	c.HTML(200,"index.html", nil)
}

func LoginPage(c *gin.Context){
	c.HTML(200,"login.html",nil)
}

func LoginPost(c *gin.Context){
	user := c.PostForm("username")
	pass := c.PostForm("password")

	var userModel models.User

	db.First(&userModel,"username = ?", user)

	err := bcrypt.CompareHashAndPassword([]byte(userModel.Password),[]byte(pass))
	if err != nil {
		c.HTML(200,"login.html",gin.H{"error": err})
		return
	}
	// Password successful, redirect to dashboard
	c.Redirect(200,"/dashboard")
}

func RegistrationPage(c *gin.Context){
	c.HTML(200,"register.html",nil)
}

func AboutPage(c *gin.Context){
	c.HTML(200,"about.html",nil)
}

func ContactPage(c *gin.Context){
	c.HTML(200,"contact.html",nil)
}

func ContactPost(c *gin.Context){

}