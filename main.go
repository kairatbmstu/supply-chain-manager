package main

import (
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

	router.GET("/", controller.IndexControllerInstance.GetIndexPage)
	router.GET("/login", controller.LoginControllerInstance.GetLoginPage)
	router.POST("/login", controller.LoginControllerInstance.PostLogin)
	router.GET("/register/resident_type", controller.RegistrationControllerInstance.GetRegistrationType)
	router.POST("/register/resident_type", controller.RegistrationControllerInstance.PostRegistrationType)
	router.GET("/register/resident_details", controller.RegistrationControllerInstance.GetResidentDetails)
	router.POST("/register/resident_details", controller.RegistrationControllerInstance.PostResidentDetails)

	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	router.Run(":8080")
}

type ApplicationContext struct {
}
