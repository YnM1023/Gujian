package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "marinette/gujian/proto/calculation"
	rpcCal "marinette/gujian/rpc/calculation"

	"google.golang.org/grpc"
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

	pb.RegisterCalculateServer(grpcServer, &rpcCal.Service{})

	grpcServer.Serve(conn)
}
