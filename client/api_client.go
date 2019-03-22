package main

import (
	"log"
	"context"
	"google.golang.org/grpc"
	pb "gin-go-rpc/proto"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

const RPCADDRESS = "127.0.0.1:9091"
const GINPORT = ":9098"

func main(){
	conn,err := grpc.Dial(RPCADDRESS,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect:%v",err)
	}
	defer conn.Close()
	pbc := pb.NewGreeterClient(conn)

	/*name := "world"

	if len(os.Args) > 1{
		name = os.Args[1]
	}*/

	gg := gin.Default()
	gg.GET("/:name",func(c *gin.Context){
		name := c.Param("name")

		req := &pb.HelloRequest{Name:name}
		r,err := pbc.SayHello(context.Background(),req)
		if err != nil{
			log.Fatalf("could not gree:%v",err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		log.Printf("Greeting:%s",r.Message)
		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(r.Message),
		})
	})

	// Run http server
	if err := gg.Run(GINPORT); err != nil {
		log.Fatalf("could not run gin server: %v", err)
	}
}
