package rpc

type Server struct {
	opts     *ServerOptions
	services map[string]Service
}
