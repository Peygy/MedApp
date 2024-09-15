package consumer

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/peygy/medapp/internal/pkg/rabbitmq"
	"github.com/peygy/medapp/internal/services/health_service/internal/services"
)

type IConsumer interface {
	ProcessMessages(ctx context.Context) error
}

type consumer struct {
	rMQ rabbitmq.IRabbitMQServer
	hs  services.IHealthService
	log logger.ILogger

	healthDataChan chan json.RawMessage
	wg             sync.WaitGroup
}

func NewConsumer(rMQ rabbitmq.IRabbitMQServer, hs services.IHealthService, log logger.ILogger) IConsumer {
	log.Info("Consumer created")
	h := &consumer{
		rMQ:            rMQ,
		hs:             hs,
		log:            log,
		healthDataChan: make(chan json.RawMessage),
	}

	if err := h.rMQ.ConsumeMessages(h.handleMessage); err != nil {
		log.Fatalf("Failed to start consuming messages: %v", err)
	}

	return h
}

func (h *consumer) handleMessage(message json.RawMessage) {
	h.healthDataChan <- message
}

func (h *consumer) ProcessMessages(ctx context.Context) error {
	h.wg.Add(1)
	defer h.wg.Done()

	for {
		select {
		case msg := <-h.healthDataChan:
			var messageData struct {
				UserId string `json:"userId"`
			}

			if err := json.Unmarshal(msg, &messageData); err != nil {
				h.log.Errorf("Failed to unmarshal message: %v", err)
				continue
			}

			if err := h.hs.InsertHealthData(messageData.UserId); err != nil {
				h.log.Errorf("Failed to insert health data: %v", err)
			}

		case <-ctx.Done():
			h.log.Info("Stopping message processing")
			return nil
		}
	}
}
