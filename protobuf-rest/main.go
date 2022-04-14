package main

import (
	"context"
	"log"
	"net/http"

	"github.com/amsokol/examples.go/protobuf-rest/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Server struct {
	// protos.UnimplementedGreeterServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, in *protos.HelloRequest) (*protos.HelloReply, error) {
	return &protos.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	mux := runtime.NewServeMux()
	// Register Greeter

	if err := protos.RegisterGreeterHandlerServer(context.Background(), mux, NewServer()); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	log.Fatalln(gwServer.ListenAndServe())
}
