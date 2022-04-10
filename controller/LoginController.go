package controller

import (
	"log"
	"net/http"
	"strings"

	"example.com/m/v2/dto"
	"example.com/m/v2/flash"
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
	msg, _ := c.Cookie("ErrorFlash")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"ErrorFlash": msg,
	})
}

func (l LoginController) PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	log.Println(" POST /login ")
	log.Println("sessionId: ", session.ID())
	var loginForm forms.LoginForm

	if c.ShouldBind(&loginForm) == nil {
		log.Println(loginForm)
		if strings.EqualFold(loginForm.Username, "admin@mail.ru") && strings.EqualFold(loginForm.Password, "123456") {
			userDto := dto.UserDTO{
				Email:    loginForm.Username,
				Username: loginForm.Username,
				Password: loginForm.Password,
			}

			session.Set("user", userDto)
			log.Println("session : ", session)
			log.Println("session.Flashes : ", session.Flashes())
			err := session.Save()
			if err != nil {
				log.Println("error : ", err)
				c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			}

			c.Redirect(http.StatusFound, "/")
		} else {
			flash.SetErrorFlash(c, "Не верный логин или пароль")
			c.Redirect(303, "/login")
		}

	}
}

func (l LoginController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(303, "/login")
}
