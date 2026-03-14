package main

import (
	"log"
	"time"

	"github.com/aecsar/go-proto/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	book := &pb.AddressBook{
		People: []*pb.Person{
			{
				Id:    1,
				Name:  "Alice",
				Email: "alice@example.com",
				LastUpdated: &timestamppb.Timestamp{
					Seconds: time.Now().Unix(),
				},
				Phones: []*pb.Person_PhoneNumber{
					{
						Number: "555-1212", Type: pb.PhoneType_PHONE_TYPE_MOBILE,
					},
				},
			},
		},
	}

	marshalled, err := proto.Marshal(book)
	if err != nil {
		log.Fatalf("error marshaling message: %v", err)
	}

	log.Printf("marshalled: %v", marshalled)

	unmarshalled := &pb.AddressBook{}

	if err = proto.Unmarshal(marshalled, unmarshalled); err != nil {
		log.Fatalf("error unmarshaling message: %v", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)
}
