package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"GinWebService/services/register/handler"
	register "GinWebService/services/register/proto/register"
	"github.com/micro/go-micro/registry/consul"
	"GinWebService/services/register/model"
)

func main() {
	consulReg:=consul.NewRegistry()

	model.InitRedis()
	model.InitDb()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.register"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":9982"),
	)

	// Initialise service
	service.Init()


	// Register Handler
	register.RegisterRegisterHandler(service.Server(), new(handler.Register))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
