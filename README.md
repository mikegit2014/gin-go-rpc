使用gin,grpc,提供对外的rpc服务.gin的原因，对外提供http访问方式，内部使用grpc提供服务

1.必须先启动grpc服务 go run server.go
2.启动gin服务 cd client , go run api_client.go

客户端代码详见：
client/api_client.go(使用gin)
client/client.go

GET请求：http://127.0.0.1:9098/world
