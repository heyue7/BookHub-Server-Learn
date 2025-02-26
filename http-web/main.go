package main

import (
	"flag"
	"github.com/heyue7/BookHub-Server-Learn/http-web/config"
	"github.com/heyue7/BookHub-Server-Learn/http-web/router"
	pb_sys_v1 "github.com/heyue7/BookHub-Server-Learn/proto-pb/system/v1"
	"github.com/heyue7/BookHub-Server-Learn/service/srv"
	"log"
)

var yamlFile = flag.String("yaml", ".yaml", "yaml file")

func init() {
	srv.FlagInit()

	err := config.Init(*yamlFile)
	if err != nil {
		log.Fatalf("config init error: %v", err)
		return
	}
	
	err = pb_sys_v1.ClientInit(config.Conf.Services.GrpcSystem)
	if err != nil {
		log.Fatalf("grpc system client init error: %v", err)
		return
	}
}

func main() {
	s := srv.NewServer()
	router.Routes(s.Group("/api/v1"))
	s.Run(config.Conf.Server.Name, config.Conf.Server.Addr)
}
