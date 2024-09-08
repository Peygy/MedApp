package main

import (
	"github.com/peygy/medapp/internal/pkg/context"
	"github.com/peygy/medapp/internal/pkg/database/postgres"
	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/peygy/medapp/internal/services/health_service/config"
	"github.com/peygy/medapp/internal/services/health_service/internal/data"
	grpcConn "github.com/peygy/medapp/internal/services/health_service/internal/grpc"
	"github.com/peygy/medapp/internal/services/health_service/internal/services"
	"github.com/peygy/medapp/internal/services/health_service/server"
	"go.uber.org/fx"
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

				services.NewHealthService,
			),
			fx.Invoke(data.InitDatabaseSchema),
			fx.Invoke(grpcConn.InitHealhGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
