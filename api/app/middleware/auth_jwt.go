package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/conan080262/go-tester-api/configs"
	"github.com/conan080262/go-tester-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AuthJWT() gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {

		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"msg":   "ไม่พบ Auth",
			})
			defer c.AbortWithStatus(http.StatusUnauthorized)
		}

		tokenHeader := c.GetHeader("Authorization")
		accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
		// fmt.Println(accessToken)
		jwtSecretKey := os.Getenv("JWT_SECRET")

		token, _ := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"msg":   "การยืนยันตัวตนผิดพลาด",
			})
			defer c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// global value result
			claims := token.Claims.(jwt.MapClaims)
			objectId, _ := primitive.ObjectIDFromHex(claims["id"].(string))

			opt := options.FindOne().SetProjection(bson.M{"cartlist": 0})

			if _, ok := claims["role"]; ok {
				if claims["role"].(string) == "1111" {
					var user models.User
					coll := configs.GetMongoCollection("users")
					_ = coll.FindOne(configs.Ctx, bson.M{"_id": objectId}, opt).Decode(&user)
					c.Set("user", user)
					c.Next()
				} else if claims["role"].(string) == "11" || claims["role"].(string) == "1" {
					var staff models.Staff
					coll := configs.GetMongoCollection("staffs")
					_ = coll.FindOne(configs.Ctx, bson.M{"_id": objectId}).Decode(&staff)
					c.Set("user", staff)
					c.Next()
				}

			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
					"msg":   "ไม่พบ role",
				})
				defer c.AbortWithStatus(http.StatusUnauthorized)
			}

		}

	})
}
