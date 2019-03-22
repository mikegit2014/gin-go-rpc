package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "gin-go-rpc/proto"
)

type server struct{}

func (s *server) SayHello(ctx context.Context,in *pb.HelloRequest) (*pb.HelloReply,error){
	log.Printf("Receive Message")
	return &pb.HelloReply{Message:"Hello "+in.Name},nil
}

const PORT = "9091"

func main(){

	lis,err := net.Listen("tcp",":"+PORT)
	if err != nil{
		log.Fatalf("net.Listen err:%v",err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})
	s.Serve(lis)
}
