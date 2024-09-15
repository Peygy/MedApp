package main

import (
	"go.uber.org/fx"

	"github.com/peygy/medapp/internal/pkg/context"
	"github.com/peygy/medapp/internal/pkg/database/postgres"
	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/peygy/medapp/internal/pkg/rabbitmq"

	"github.com/peygy/medapp/internal/services/auth_service/config"
	"github.com/peygy/medapp/internal/services/auth_service/internal/data"
	grpcConn "github.com/peygy/medapp/internal/services/auth_service/internal/grpc"
	"github.com/peygy/medapp/internal/services/auth_service/internal/managers"
	"github.com/peygy/medapp/internal/services/auth_service/server"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewAuthConfig,
				logger.NewLogger,
				context.NewContext,
				grpc.NewGrpcServer,
				postgres.NewDatabaseConnection,
				rabbitmq.NewRabbitMQConnection,

				managers.NewPasswordManager,
				managers.NewRoleManager,
				managers.NewUserManager,
			),
			fx.Invoke(data.InitDatabaseSchema),
			fx.Invoke(grpcConn.InitAuthGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
