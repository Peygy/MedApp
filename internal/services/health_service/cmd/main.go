package main

import (
	"go.uber.org/fx"

	"github.com/peygy/medapp/internal/pkg/context"
	"github.com/peygy/medapp/internal/pkg/database/postgres"
	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/peygy/medapp/internal/pkg/rabbitmq"

	"github.com/peygy/medapp/internal/services/health_service/config"
	"github.com/peygy/medapp/internal/services/health_service/internal/consumer"
	"github.com/peygy/medapp/internal/services/health_service/internal/data"
	grpcConn "github.com/peygy/medapp/internal/services/health_service/internal/grpc"
	"github.com/peygy/medapp/internal/services/health_service/internal/services"
	"github.com/peygy/medapp/internal/services/health_service/server"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewHealthConfig,
				logger.NewLogger,
				context.NewContext,
				grpc.NewGrpcServer,
				postgres.NewDatabaseConnection,

				services.NewHealthService,
				rabbitmq.NewRabbitMQConnection,
				consumer.NewConsumer,
			),
			fx.Invoke(data.InitDatabaseSchema),
			fx.Invoke(grpcConn.InitHealhGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
