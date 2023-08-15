package rpc

import "context"

// Service defines a generic implementation interface for a specific Service
type Service interface {
	Register(string, Handler)
	Serve(*ServerOptions)
	Close()
}

type service struct {
	svr         interface{}        // server
	ctx         context.Context    // Each service is managed in one context
	cancel      context.CancelFunc // controller of context
	serviceName string             // service name
	handlers    map[string]Handler
	opts        *ServerOptions // parameter options
}
