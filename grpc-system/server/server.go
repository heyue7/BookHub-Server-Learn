package server

import (
	pb_sys_v1 "github.com/heyue7/BookHub-Server-Learn/proto-pb/system/v1"
)

type Server struct {
	pb_sys_v1.UnimplementedGetterServer
}
