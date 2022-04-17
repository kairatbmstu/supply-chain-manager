package flash

import (
	"github.com/gin-gonic/gin"
)

func SetErrorFlash(c *gin.Context, msg string) {
	c.SetCookie("ErrorFlash", msg, 1, "/", "localhost", true, true)
}

func SetFlash(c *gin.Context, fieldName, msg string) {
	c.SetCookie(fieldName, msg, 1, "/", "localhost", true, true)
}
