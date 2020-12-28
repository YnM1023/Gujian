package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "The backend server port")
)

func main() {
	flag.Parse()

	conn, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	var opt []grpc.ServerOption
	grpcServer := grpc.NewServer(opt...)
	grpcServer.Serve(conn)
}
