package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "goweb/chapter5/protobuf/chapter5/protobuf"
	"log"
	"net"
)

type ProgrammerServiceServer struct {
	pb.UnimplementedProgrammerServiceServer
}

func (p *ProgrammerServiceServer) GetProgrammerInfo(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "shirdon" {
		resp = &pb.Response{
			Uid:      6,
			Username: name,
			Job:      "CTO",
			GoodAt:   []string{"Go", "Java", "PHP", "Python"},
		}
	}
	err = nil
	return
}

func main() {
	port := ":8078"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error:%v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()
	pb.RegisterProgrammerServiceServer(s, &ProgrammerServiceServer{})
	s.Serve(l)
}
