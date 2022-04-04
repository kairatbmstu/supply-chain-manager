package controller

import (
	"log"
	"net/http"
	"strings"

	"example.com/m/v2/dto"
	"example.com/m/v2/forms"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var LoginControllerInstance LoginController

type LoginController struct {
}

func (l LoginController) GetLoginPage(c *gin.Context) {
	session := sessions.Default(c)
	log.Println(" GET /login ")
	log.Println("sessionId: ", session.ID())
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (l LoginController) PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	log.Println(" POST /login ")
	log.Println("sessionId: ", session.ID())
	var loginForm forms.LoginForm

	if c.ShouldBind(&loginForm) == nil {
		log.Println(loginForm)
		if strings.EqualFold(loginForm.Username, "admin") && strings.EqualFold(loginForm.Password, "123456") {

			userDto := dto.UserDTO{
				Email:    loginForm.Username,
				Username: loginForm.Username,
				Password: loginForm.Password,
			}

			session.Set("user", userDto)
			err := session.Save()
			if err != nil {
				log.Println("error : ", err)
				c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			}

			c.Redirect(http.StatusFound, "/")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		}

	}
}
