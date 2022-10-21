package server

import "time"

type serverOptions struct {
	timeout        time.Duration
	maxConnections int
}

func (o serverOptions) getTimeout() time.Duration {
	if o.timeout == 0 {
		return 10 * time.Second
	}
	return o.timeout
}

func (o serverOptions) getMaxConnections() int {
	if o.maxConnections == 0 {
		return 5
	}
	return o.maxConnections
}
