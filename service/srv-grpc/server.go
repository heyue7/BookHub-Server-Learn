package srv_grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

type Server struct {
	*grpc.Server

	conf Config
	lis  *net.Listener
}

func NewServer(conf Config) *Server {
	return &Server{
		conf:   conf,
		Server: grpc.NewServer(),
		lis:    new(net.Listener),
	}
}

func (s *Server) Serve() (err error) {
	addr := s.conf.Addr
	if addr == "" {
		addr = "0.0.0.0:0"
	}
	if !strings.Contains(addr, ":") {
		addr += ":0"
	}

	*s.lis, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()

		log.Printf("%s started. Press any key to stop.\n", s.conf.ServerName)
		var str string
		fmt.Scanf("%s", &str)
		s.Server.Stop()
	}()

	if err = s.Server.Serve(*s.lis); err != nil {
		return
	}

	return
}
