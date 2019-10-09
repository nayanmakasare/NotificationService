package main

import (
	"context"
	NotificationService "github.com/nayanmakasare/NotificationService/proto"
	"github.com/streadway/amqp"
)

type NotificationServiceSubscriber struct {
	rabbitChannel *amqp.Channel
}

func(nss *NotificationServiceSubscriber) NotificationEvent(ctx context.Context, messageToPublish *NotificationService.MessageToPublish) error{
	return nss.rabbitChannel.Publish(
		messageToPublish.ExchangeName, //exchange
		messageToPublish.RoutingKeyName, //routing key
		false, //mandatory
		false, //immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body: messageToPublish.MessageTosend,
		})
}


