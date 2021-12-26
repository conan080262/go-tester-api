package main

import (
	"os"

	"github.com/conan080262/go-tester-api/configs"
	"github.com/conan080262/go-tester-api/models"
	routesv1 "github.com/conan080262/go-tester-api/routes/v1"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	router := SetupRouter()
	router.Run(":" + os.Getenv("GO_PORT"))
	// gin.SetMode(gin.ReleaseMode)

}

func SetupRouter() *gin.Engine {
	//Load .env
	godotenv.Load(".env")
	gin.SetMode(os.Getenv("GIN_MODE"))
	// gin.SetMode(gin.ReleaseMode)
	configs.MongoConnection()
	models.CreateUserManyIndexUnique(configs.Ctx)
	models.CreateStaffManyIndexUnique(configs.Ctx)

	var maxBytes int64 = 1024 * 1024 * 10 // 10MB
	router := gin.Default()
	router.Use(limits.RequestSizeLimiter(maxBytes))
	router.Static("/go/api/v1/public/images", "./stores")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "paaaa",
		})
	})

	apiV1 := router.Group("/go/api/v1")
	routesv1.InitHomeRoutes(apiV1)
	routesv1.InitUserRoutes(apiV1)
	routesv1.InitStaffRoutes(apiV1)
	routesv1.InitProductRoutes(apiV1)
	routesv1.InitOrderRoutes(apiV1)
	return router
}
