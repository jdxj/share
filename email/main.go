package main

import (
	"log"

	//"github.com/asim/go-micro/plugins/transport/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/broker"

	"github.com/jdxj/share/email/handler"
	email "github.com/jdxj/share/email/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("email"),
		//micro.Transport(grpc.NewTransport()),
	)
	service.Init()
	err := email.RegisterEmailHandler(service.Server(), &handler.Email{})
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	//err = service.Options().Broker.Init()
	//if err != nil {
	//	log.Fatalf("Init: %s\n", err)
	//}
	//err = service.Options().Broker.Connect()
	//if err != nil {
	//	log.Fatalf("Connect: %s\n", err)
	//}
	//_, err = service.Options().Broker.Subscribe("testbroker", Handle)
	//if err != nil {
	//	log.Fatalf("Subscribe: %s\n", err)
	//}

	err = service.Run()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

func Handle(event broker.Event) error {
	log.Printf("header: %v\n", event.Message().Header)
	return nil
}
