package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/dto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var IndexControllerInstance IndexController

type IndexController struct {
}

func (l IndexController) GetIndexPage(c *gin.Context) {
	session := sessions.Default(c)
	log.Println(" user ", session.Get("user"))
	log.Println(" Username ", session.Get("Username"))
	log.Println(" Password ", session.Get("Password"))

	log.Println(" GET / ")
	log.Println("sessionId: ", session.ID())

	if session.Get("user") != nil {
		userdto := session.Get("user").(dto.UserDTO)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"email":    userdto.Email,
			"username": userdto.Username,
			"password": userdto.Password,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}

}
