package pb_sys_v1

import (
	srv_grpc "github.com/heyue7/BookHub-Server-Learn/service/srv-grpc"
)

var __client GetterClient

func ClientInit(addr string) error {
	cc, err := srv_grpc.NewClient(addr)
	if err != nil {
		return err
	}

	__client = NewGetterClient(cc)
	return nil
}

func Client() GetterClient {
	return __client
}
