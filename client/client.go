package main

import(
	"log"
	"os"
	"context"
	"google.golang.org/grpc"
	pb "gin-go-rpc/proto"
)

const ADDRESS = "127.0.0.1:9091"

func main(){
	conn,err := grpc.Dial(ADDRESS,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect:%v",err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := "world"

	if len(os.Args) > 1{
		name = os.Args[1]
	}

	r,err := c.SayHello(context.Background(),&pb.HelloRequest{Name:name})
	if err != nil{
		log.Fatalf("could not greet:%v",err)
	}
	log.Printf("Greeting:%s",r.Message)
}
