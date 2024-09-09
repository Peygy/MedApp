package rabbitmq

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *rabbitMQServer) PublishMessage(message interface{}) error {
	messageBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = r.channel.Publish(
		"",          // exchange
		r.queueName, // routing key (queue name)
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBody,
		},
	)
	if err != nil {
		return fmt.Errorf("rabbitmq: failed to publish a message: %v", err)
	}
	r.log.Infof("Message published: %s", string(messageBody))
	return nil
}

func (r *rabbitMQServer) ConsumeMessages(handler func(message json.RawMessage)) error {
	messages, err := r.channel.Consume(
		r.queueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return fmt.Errorf("rabbitmq: failed to register a consumer: %v", err)
	}

	go func() {
		for msg := range messages {
			handler(msg.Body)
		}
	}()

	return nil
}
