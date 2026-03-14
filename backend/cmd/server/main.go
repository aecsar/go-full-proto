package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	pb "github.com/aecsar/go-proto/gen/pb/aecsar/go_proto/v1"
	pbconnect "github.com/aecsar/go-proto/gen/pb/aecsar/go_proto/v1/go_protov1connect"
)

type GreetServer struct{}

func (s *GreetServer) Greet(_ context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	res := &pb.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Name),
	}

	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()

	path, handler := pbconnect.NewGreetServiceHandler(greeter, connect.WithInterceptors(validate.NewInterceptor()))

	mux.Handle(path, handler)

	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true) // Use h2c so we can serve HTTP/2 without TLS.

	srv := http.Server{
		Addr:      ":3000",
		Handler:   mux,
		Protocols: p,
	}

	fmt.Println("server listening on port 3000")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting server : %v", err)
	}
}
