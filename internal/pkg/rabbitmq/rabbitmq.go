package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/peygy/medapp/internal/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IRabbitMQServer interface {
	PublishMessage(message interface{}) error
	ConsumeMessages(handler func(message json.RawMessage)) error

	Run(ctx context.Context) error
}

type rabbitMQServer struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string

	log logger.ILogger
}

type RabbitMQConfig struct {
	QueueName string `yaml:"queueName"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
}

// NewRabbitMQConnection создает новое подключение к RabbitMQ
func NewRabbitMQConnection(cfg *RabbitMQConfig, log logger.ILogger) (IRabbitMQServer, error) {
	rabbitConnStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, cfg.Password, cfg.Host, cfg.Port)

	conn, err := amqp.Dial(rabbitConnStr)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error creating RabbitMQ channel: %v", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		cfg.QueueName, // queue name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq: failed to declare a queue: %v", err)
	}

	log.Info("Connected to RabbitMQ")
	return &rabbitMQServer{
		conn:      conn,
		channel:   ch,
		queueName: cfg.QueueName,
		log:       log,
	}, nil
}

// Run обрабатывает завершение соединения
func (r *rabbitMQServer) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		r.log.Info("Shutting down RabbitMQ channel")
		r.channel.Close()
		r.conn.Close()
		r.log.Info("RabbitMQ channel exited properly")
	}()

	return nil
}
