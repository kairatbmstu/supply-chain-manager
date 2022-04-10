package middleware

import (
	"log"
	"strings"

	sessions "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckAuthentification(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	log.Println("c.Request.URL.RawPath : ", c.Request.URL.String())

	if strings.Contains(c.Request.URL.String(), "/login") ||
		strings.Contains(c.Request.URL.String(), "/logout") ||
		strings.Contains(c.Request.URL.String(), "/register") ||
		strings.Contains(c.Request.URL.String(), "/web/js") ||
		strings.Contains(c.Request.URL.String(), "/web/css") {

	} else {
		if user == nil {
			c.Redirect(303, "/login")
		}
	}

}
