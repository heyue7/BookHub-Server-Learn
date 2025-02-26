package srv

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type iController interface {
	DoHandle(c *gin.Context) *Response
}

func Handler(controller iController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := controller.DoHandle(ctx)

		if resp == nil {
			return
		}

		switch strings.ToUpper(ctx.Request.Method) {
		case "POST", "GET":
		default:
			ctx.JSON(200, resp)
			return
		}

		switch strings.ToLower(ctx.Request.Header.Get("Content-Type")) {
		case "multipart/form-data":
			ctx.JSON(200, resp)
			return
		}

		// TODO: 加密

		ctx.JSON(200, resp)
	}
}
