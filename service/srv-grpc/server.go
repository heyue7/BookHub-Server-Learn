package srv_grpc

import (
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	conf Config

	*grpc.Server

	lis *net.Listener
}

func New(conf Config) *Server {
	return &Server{
		conf:   conf,
		Server: grpc.NewServer(),
	}
}

func (s *Server) Serve() (err error) {
	*s.lis, err = net.Listen("tcp", s.conf.Addr)
	if err != nil {
		return err
	}

	if err = s.Server.Serve(*s.lis); err != nil {
		return err
	}

	return
}
