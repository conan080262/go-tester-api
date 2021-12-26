package routesv1

import (
	productcontroller "github.com/conan080262/go-tester-api/controllers/product"
	middlewares "github.com/conan080262/go-tester-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/products")
	routerGroup.POST("/auth/insert", middlewares.AuthJWT(), productcontroller.InsertProduct)
	routerGroup.PUT("/auth/update", middlewares.AuthJWT(), productcontroller.UpdateProduct)
	routerGroup.DELETE("/auth/delete", middlewares.AuthJWT(), productcontroller.DeleteProduct)
	routerGroup.GET("/fetchdata", productcontroller.FetchProduct)
	routerGroup.GET("/get", productcontroller.GetProduct)
	routerGroup.POST("/addtocart", productcontroller.AddToUserCart)
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VERSION": "1.0.0",
		})
	})
}
