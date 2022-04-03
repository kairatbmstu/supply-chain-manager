package main

import (
	"net/http"

	"example.com/m/v2/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Static("/css", "web/css")
	router.Static("/js", "web/js")
	router.LoadHTMLGlob("web/templates/**/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Users",
		})
	})
	router.GET("/register/resident_type", controller.RegistrationControllerInstance.GetRegistrationType)
	router.POST("/register/resident_type", controller.RegistrationControllerInstance.PostRegistrationType)
	router.GET("/register/resident_details", controller.RegistrationControllerInstance.GetResidentDetails)
	router.POST("/register/resident_details", controller.RegistrationControllerInstance.PostResidentDetails)

	router.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})

	router.GET("/hello2", func(c *gin.Context) {
		session := sessions.Default(c)

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})

	router.Run(":8080")
}

type ApplicationContext struct {
}
