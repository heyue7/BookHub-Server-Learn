package server

import (
	"context"
	"github.com/heyue7/BookHub-Server-Learn/grpc-system/model"
	pb_sys_v1 "github.com/heyue7/BookHub-Server-Learn/proto-pb/system/v1"
	"github.com/heyue7/BookHub-Server-Learn/service/srv"
	"google.golang.org/grpc/status"
)

func (s Server) PermissionGet(ctx context.Context, req *pb_sys_v1.PermissionGet_Request) (resp *pb_sys_v1.PermissionGet_Response, err error) {
	resp = &pb_sys_v1.PermissionGet_Response{}

	if req.GetCode() == "" {
		err = status.Error(1001, "code is required")
		return
	}

	db := srv.DB()

	if req.GetCode() != "" {
		db.Where("code = ?", req.GetCode())
	} else {
		err = status.Error(1001, "code is required")
		return
	}

	var m model.PermissionModel

	if _, err = db.Get(&m); err != nil {
		err = status.Error(1002, "permission not found, "+err.Error())
		return
	}

	return
}
