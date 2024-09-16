package config

import (
	"github.com/peygy/medapp/internal/pkg/config"
	"github.com/peygy/medapp/internal/pkg/database/postgres"
	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/rabbitmq"
)

const configPath = "./config/config.dev.yml"

type HealthConfig struct {
	GrpcServer     *grpc.GrpcServerConfig   `yaml:"grpc-server"`
	DatabaseConfig *postgres.PostgresConfig `yaml:"database"`
	RabbitMQConfig *rabbitmq.RabbitMQConfig `yaml:"rabbitmq"`
}

func NewHealthConfig() (*HealthConfig, *grpc.GrpcServerConfig, *postgres.PostgresConfig, *rabbitmq.RabbitMQConfig, error) {
	cfg, err := config.NewConfig[HealthConfig](configPath)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return cfg, cfg.GrpcServer, cfg.DatabaseConfig, cfg.RabbitMQConfig, nil
}
