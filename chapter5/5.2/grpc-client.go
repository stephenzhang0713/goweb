package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "goweb/chapter5/protobuf/chapter5/protobuf"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8078", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error:%v\n", err)
	}
	defer conn.Close()

	client := pb.NewProgrammerServiceClient(conn)

	req := new(pb.Request)
	req.Name = "shirdon"
	resp, err := client.GetProgrammerInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error:%v\n", err)
	}
	fmt.Printf("Received:%v\n", resp)
}
