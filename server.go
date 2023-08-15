package rpc

type Server struct {
	opts     *ServerOptions
	services map[string]Service
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
