package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseJWT(c *gin.Context) (Id int) {
	tokenString, _ := c.Request.Cookie("id")
	value := tokenString.Value
	id, err := strconv.Atoi(value)
	if err != nil {
		NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	return id
}
