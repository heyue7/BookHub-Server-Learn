package router

import (
	"github.com/gin-gonic/gin"
	permission_controller "github.com/heyue7/BookHub-Server-Learn/http-web/controller/permission-controller"
	"github.com/heyue7/BookHub-Server-Learn/service/srv"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/permission/get", srv.Handler(permission_controller.PermissionGet{}))
}
