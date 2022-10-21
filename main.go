package main

import (
	"fmt"
	"functional_options/pkg/server"
	"time"
)

func main() {
	options := server.Options(
		server.Timeout(10*time.Second),
		server.MaxConnections(10))
	srv := server.NewServer(":8080", options)

	fmt.Printf("%+v", srv)
}
