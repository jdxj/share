package main

import (
	"context"
	"log"
	"time"

	"github.com/asim/go-micro/v3/errors"

	//"github.com/asim/go-micro/plugins/transport/grpc/v3"
	email "github.com/jdxj/share/email/proto"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/broker"
)

func main() {
	service := micro.NewService(
		micro.RegisterTTL(time.Hour),
		//micro.Transport(grpc.NewTransport()),
	)
	service.Init()

	//err := service.Options().Broker.Init()
	//if err != nil {
	//	log.Fatalf("Init: %s\n", err)
	//}
	//err = service.Options().Broker.Connect()
	//if err != nil {
	//	log.Fatalf("Connect: %s\n", err)
	//}
	//_, err = service.Options().Broker.Subscribe("test broker", Handle)
	//if err != nil {
	//	log.Fatalf("Subscribe: %s\n", err)
	//}

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	err := service.Options().Broker.Publish("testbroker", &broker.Message{
	//		Header: map[string]string{
	//			"abc": "def",
	//		},
	//		Body: []byte("hello world"),
	//	})
	//	if err != nil {
	//		log.Fatalf("Publish: %s\n", err)
	//	}
	//	log.Printf("pub ok!\n")
	//}()

	e := email.NewEmailService("email", service.Client())
	//log.Printf("%d\n", time.Now().Unix())

	ctx, cancel := context.WithTimeout(context.TODO(), time.Hour)
	defer cancel()

	resp, err := e.Send(ctx, &email.RequestEmail{})
	if err != nil {
		e := errors.Parse(err.Error())
		log.Printf("%#v\n", e)

		log.Printf("%d\n", time.Now().Unix())
		log.Fatalf("Send: %s\n", err.Error())
	}
	msg := &email.ResponseEmail{}
	err = resp.RecvMsg(msg)
	if err != nil {
		log.Fatalf("RecvMsg: %s\n", err)
	}
	log.Printf("%#v\n", msg)

	//log.Printf("%s\n", resp)
	if err := service.Run(); err != nil {
		log.Fatalf("Run: %s\n", err)
	}
}

func Handle(event broker.Event) error {
	log.Printf("header: %v\n", event.Message().Header)
	return nil
}
