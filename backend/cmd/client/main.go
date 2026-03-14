package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/aecsar/go-proto/gen/pb/aecsar/go_proto/v1"
	pbconnect "github.com/aecsar/go-proto/gen/pb/aecsar/go_proto/v1/go_protov1connect"
)

func main() {
	client := pbconnect.NewGreetServiceClient(http.DefaultClient, "http://localhost:3000")

	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		Name: "Ousseynou",
	})

	if err != nil {
		log.Fatalf("error sending greeting req: %v", err)
	}

	fmt.Printf("response: %v\n", res.Greeting)

	invalid, err := client.Greet(context.Background(), &pb.GreetRequest{
		Name: "",
	})

	if err != nil {
		log.Fatalf("error sending greeting req: %v", err)
	}

	fmt.Printf("invalid res : %v\n", invalid.Greeting)
}
