// Code generated by hertz generator.

package main

import (
	"grpc/biz/rpc"
	"grpc/config"
	"grpc/dal"
	"log"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Init() {
	rpc.InitConvertRPC()
	log.Println("successfully register rpc client")
	config.ReadConfig()
	dal.Init()

	log.Println("successfully running...")
	klog.SetLevel(klog.LevelDebug)
	// rpc.Init()
	// select {
	// case <-rpc.ClientReady:
	// 	log.Println("gRPC client successfully initialized")
	// case <-time.After(10 * time.Second):
	// 	log.Fatal("Failed to initialize gRPC client within timeout period")
	// }

}

func main() {
	Init()

	h := server.Default(server.WithHostPorts("localhost:8888"),
		server.WithMaxRequestBodySize(419430400))

	register(h)
	h.Spin()
}
