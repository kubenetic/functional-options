package server

import "time"

type Option func(s *serverOptions)

func Options(opts ...Option) Option {
	return func(s *serverOptions) {
		for _, opt := range opts {
			opt(s)
		}
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *serverOptions) {
		s.timeout = timeout
	}
}

func MaxConnections(connections int) Option {
	return func(s *serverOptions) {
		s.maxConnections = connections
	}
}
