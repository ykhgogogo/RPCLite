package RPCLite

import (
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	opts     *ServerOptions
	services map[string]Service
	closing  bool
}

// 创建Server
func NewServer(opt ...ServerOption) *Server {

	s := &Server{
		opts:     &ServerOptions{},
		services: make(map[string]Service),
	}

	for _, o := range opt {
		o(s.opts)
	}

	return s
}

// 服务运行
func (s *Server) Serve() {
	for _, service := range s.services {
		go service.Serve(s.opts)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSEGV)
	<-ch

	s.Close()
}

func (s *Server) Close() {

}
