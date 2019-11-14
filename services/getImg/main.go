package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"GinWebService/services/getImg/handler"
	getImg "GinWebService/services/getImg/proto/getImg"
	"github.com/micro/go-micro/registry/consul"
	"GinWebService/services/getImg/model"
)

func main() {
	consulReg:=consul.NewRegistry()
	// New Service
	model.InitRedis()
	service := micro.NewService(
		micro.Name("go.micro.srv.getImg"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":9981"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getImg.RegisterGetImgHandler(service.Server(), new(handler.GetImg))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
