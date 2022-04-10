package flash

import (
	"github.com/gin-gonic/gin"
)

func SetErrorFlash(c *gin.Context, msg string) {
	c.SetCookie("ErrorFlash", msg, 1, "/", "localhost", true, true)
}
