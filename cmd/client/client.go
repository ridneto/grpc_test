package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ridneto/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %w", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "neto",
		Email: "neto@test.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %w", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "neto",
		Email: "neto@test.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %w", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the msg: %w", err)
		}

		fmt.Println("Status:", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "neto 1",
			Email: "neto1@test.com",
		},
		{
			Id:    "2",
			Name:  "neto 2",
			Email: "neto2@test.com",
		},
		{
			Id:    "3",
			Name:  "neto 3",
			Email: "neto3@test.com",
		},
		{
			Id:    "4",
			Name:  "neto 4",
			Email: "neto4@test.com",
		},
		{
			Id:    "5",
			Name:  "neto 5",
			Email: "neto5@test.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "neto 1",
			Email: "neto1@test.com",
		},
		{
			Id:    "2",
			Name:  "neto 2",
			Email: "neto2@test.com",
		},
		{
			Id:    "3",
			Name:  "neto 3",
			Email: "neto3@test.com",
		},
		{
			Id:    "4",
			Name:  "neto 4",
			Email: "neto4@test.com",
		},
		{
			Id:    "5",
			Name:  "neto 5",
			Email: "neto5@test.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data $v", err)
			}

			fmt.Printf("Receive user %v with status %v\n", res.GetUser(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
