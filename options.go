package RPCLite

import "time"

// ServerOptions defines the server serve parameters
type ServerOptions struct {
	address           string        // listening address, e.g. :( ip://127.0.0.1:8080、 dns://www.google.com)
	network           string        // network type, e.g. : tcp、udp
	protocol          string        // protocol type, e.g. : proto、json
	timeout           time.Duration // timeout
	serializationType string        // serialization type, default: proto

	selectorSvrAddr string   // service discovery server address, required when using the third-party service discovery plugin
	tracingSvrAddr  string   // tracing plugin server address, required when using the third-party tracing plugin
	tracingSpanName string   // tracing span name, required when using the third-party tracing plugin
	pluginNames     []string // plugin name
	interceptors    []interceptor.ServerInterceptor
}

type ServerOption func(*ServerOptions)
