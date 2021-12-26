package routesv1

import (
	usercontroller "github.com/conan080262/go-tester-api/controllers/user"
	middlewares "github.com/conan080262/go-tester-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")
	routerGroup.POST("/register", usercontroller.Register)
	routerGroup.POST("/login", usercontroller.Login)
	routerGroupAuth := routerGroup.Group("/auth").Use(middlewares.AuthJWT())
	routerGroupAuth.POST("/pushaddbook", usercontroller.AddToAddressBooks)
	routerGroupAuth.PATCH("/pulladdbook", usercontroller.PullUserAddr)
	routerGroupAuth.PATCH("/pullcart", usercontroller.PullUserCart)
	routerGroupAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VERSION": c.MustGet("user"),
		})
	})
}
