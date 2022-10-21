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

	for _, opt := range opts {
		opt(server)
	}

	return server
}
