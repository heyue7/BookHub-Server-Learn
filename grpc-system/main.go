package main

import (
	"flag"
	"github.com/heyue7/BookHub-Server-Learn/grpc-system/config"
	"github.com/heyue7/BookHub-Server-Learn/grpc-system/server"
	pb_sys_v1 "github.com/heyue7/BookHub-Server-Learn/proto-pb/system/v1"
	"github.com/heyue7/BookHub-Server-Learn/service/srv"
	srv_db "github.com/heyue7/BookHub-Server-Learn/service/srv-db"
	srv_grpc "github.com/heyue7/BookHub-Server-Learn/service/srv-grpc"
	"log"
)

var yamlFile = flag.String("yaml", ".yaml", "yaml file")

func init() {
	srv.FlagInit()

	if *yamlFile == "" {
		log.Fatalf("yaml file is empty")
		return
	}

	if err := config.Init(*yamlFile); err != nil {
		log.Fatalf("config init error: %v", err)
		return
	}

	if err := srv_db.Init(config.Conf.DB); err != nil {
		log.Fatalf("db init error: %v", err)
		return
	}

}

func main() {
	s := srv_grpc.New(config.Conf.Server)

	pb_sys_v1.RegisterGetterServer(s.Server, &server.Server{})

	err := s.Serve()
	if err != nil {
		log.Fatalf("grpc server serve error: %v", err)
		return
	}
}
