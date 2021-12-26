package routesv1

import (
	staffcontroller "github.com/conan080262/go-tester-api/controllers/staff"
	"github.com/gin-gonic/gin"
)

func InitStaffRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/staffs")
	routerGroup.POST("/login", staffcontroller.Login)
	routerGroup.POST("/register", staffcontroller.Register)
	routerGroup.POST("/uploadfile", staffcontroller.UploadImg)
}
