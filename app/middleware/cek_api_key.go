package middleware

import "github.com/gin-gonic/gin"

func CekAPIKey(c *gin.Context) {
	c.Next()
	return

	apiKey := c.GetHeader("X-API-KEY")
	if apiKey != "SECRET" {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	c.Next()
}
