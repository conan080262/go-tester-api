package routesv1

import (
	ordercontroller "github.com/conan080262/go-tester-api/controllers/order"
	"github.com/gin-gonic/gin"
)

func InitOrderRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/orders")
	routerGroup.POST("/auth/insert", ordercontroller.InsertOrder)
	routerGroup.GET("/auth/fetchdata", ordercontroller.FetchOrder)
	routerGroup.PATCH("/auth/addpayment", ordercontroller.UserAddPayment)
	routerGroup.PATCH("/auth/staffcarry", ordercontroller.StaffCarry)
	routerGroup.PATCH("/auth/cancelorder", ordercontroller.CancelOrder)
	routerGroup.POST("/auth/myorder", ordercontroller.MyOrder)
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VERSION": "1.0.0",
		})
	})
}
