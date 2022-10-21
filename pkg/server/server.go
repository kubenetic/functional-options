package server

import "time"

type Server struct {
	addr           string
	timeout        time.Duration
	maxConnections int
}

func NewServer(addr string, opts ...Option) *Server {
	server := &Server{
		addr: addr,
	}

	options := &serverOptions{}
	for _, opt := range opts {
		opt(options)
	}

	server.timeout = options.getTimeout()
	server.maxConnections = options.getMaxConnections()

	return server
}
