package server

import "time"

type Option func(s *Server)

func Options(opts ...Option) Option {
	return func(s *Server) {
		for _, opt := range opts {
			opt(s)
		}
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func MaxConnections(connections int) Option {
	return func(s *Server) {
		s.maxConnections = connections
	}
}
