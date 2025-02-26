package permission_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	pb_sys_v1 "github.com/heyue7/BookHub-Server-Learn/proto-pb/system/v1"
	"github.com/heyue7/BookHub-Server-Learn/service/srv"
)

type PermissionGet struct {
	Request struct {
		Code string `json:"code"`
	}
	Response struct {
		Code       string `json:"code"`
		GroupCode  string `json:"group_code"`
		Name       string `json:"name"`
		CreateTime int64  `json:"create_time"`
		UpdateTime int64  `json:"update_time"`
	}
}

func (this PermissionGet) DoHandle(ctx *gin.Context) *srv.Response {
	if err := ctx.ShouldBind(&this.Request); err != nil {
		return srv.ErrorWithValidate(err)
	}

	if this.Request.Code == "" {
		return srv.Error(7002, "Code is required")
	}

	req := &pb_sys_v1.PermissionGet_Request{Code: this.Request.Code}

	resp, err := pb_sys_v1.Client().PermissionGet(ctx, req)
	if err != nil {
		return srv.Error(7003, fmt.Sprintf("数据库查询失败：%s", err.Error()))
	}

	this.Response.Code = resp.GetPermission().GetCode()
	this.Response.GroupCode = resp.GetPermission().GetGroupCode()
	this.Response.Name = resp.GetPermission().GetName()
	this.Response.CreateTime = resp.GetPermission().GetCreateTime()
	this.Response.UpdateTime = resp.GetPermission().GetUpdateTime()

	return srv.Success(this.Response)
}
