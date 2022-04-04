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

	if session.Get("user") != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"email":    session.Get("user").(dto.UserDTO).Email,
			"username": session.Get("user").(dto.UserDTO).Username,
			"password": session.Get("user").(dto.UserDTO).Password,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}

}
