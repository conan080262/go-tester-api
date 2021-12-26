package middlewares

import (
	"net/http"

	"github.com/conan080262/go-tester-api/models"
	"github.com/gin-gonic/gin"
)

func CheckRole(check_role string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if check_role == "1" {
			if c.MustGet("user").(models.Staff).Role == check_role {
				c.Next()
			}
		} else if check_role == "11" {
			if c.MustGet("user").(models.Staff).Role == check_role {
				c.Next()
			}
		} else if check_role == "1111" {
			if c.MustGet("user").(models.User).Role == check_role {
				c.Next()
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"msg":   "Role ผิด",
		})
		defer c.AbortWithStatus(http.StatusUnauthorized)
	})
}
