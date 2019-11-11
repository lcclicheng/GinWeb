package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"GinWebService/services/getArea/handler"
	getArea "GinWebService/services/getArea/proto/getArea"
	"GinWebService/services/getArea/model"
)

func main() {
	model.InitDb()
	model.InitRedis()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getArea"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getArea.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
