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
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (l LoginController) PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	var loginForm forms.LoginForm

	if c.ShouldBind(&loginForm) == nil {
		log.Println(loginForm)
		if strings.EqualFold(loginForm.Username, "admin") && strings.EqualFold(loginForm.Password, "123456") {

			user := dto.UserDTO{
				Email:    loginForm.Username,
				Username: loginForm.Username,
				Password: loginForm.Password,
			}

			session.Set("user", user)
			session.Set("Username", loginForm.Username)
			session.Set("Password", loginForm.Password)
			session.Save()

			c.Redirect(http.StatusFound, "/")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		}

	}
}
