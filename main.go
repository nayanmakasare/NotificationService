package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/streadway/amqp"
)


func main(){
	service := grpc.NewService(
		micro.Name("NotificationService"),
		micro.Address(":50057"),
		micro.Version("1.0"),
	)
	service.Init()
	conn, err := amqp.Dial("amqp://blwmcjci:vUyxgSpStc5IgHuSi4WFdVPVnxQ8qR57@barnacle.rmq.cloudamqp.com/blwmcjci")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	rabbitSubscriber := NotificationServiceSubscriber{rabbitChannel:ch}

	//Register Subscriber
	err = micro.RegisterSubscriber("notify", service.Server(), rabbitSubscriber)
	if err != nil {
		log.Fatal(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}



