package config

import (
	"github.com/peygy/medapp/internal/pkg/config"
	"github.com/peygy/medapp/internal/pkg/gin"
	"github.com/peygy/medapp/internal/pkg/grpc"
)

const configPath = "./config/config.dev.yml"

type GraphQLConfig struct {
	Gin        *gin.GinConfig         `yaml:"gin"`
	GrpcClient *grpc.GrpcClientConfig `yaml:"grpc-client"`
}

func NewGraphQLConfig() (*GraphQLConfig, *gin.GinConfig, *grpc.GrpcClientConfig, error) {
	cfg, err := config.NewConfig[GraphQLConfig](configPath)
	if err != nil {
		return nil, nil, nil, err
	}

	return cfg, cfg.Gin, cfg.GrpcClient, nil
}
