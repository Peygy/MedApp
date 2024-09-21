package main

import (
	"go.uber.org/fx"

	"github.com/peygy/medapp/internal/pkg/context"
	"github.com/peygy/medapp/internal/pkg/database/postgres"
	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/peygy/medapp/internal/services/crew_service/config"
	"github.com/peygy/medapp/internal/services/crew_service/internal/data"
	grpcConn "github.com/peygy/medapp/internal/services/crew_service/internal/grpc"
	"github.com/peygy/medapp/internal/services/crew_service/internal/services"
	"github.com/peygy/medapp/internal/services/crew_service/server"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewCrewConfig,
				logger.NewLogger,
				context.NewContext,
				grpc.NewGrpcServer,
				postgres.NewDatabaseConnection,

				services.NewCrewService,
			),
			fx.Invoke(data.InitDatabaseSchema),
			fx.Invoke(grpcConn.InitCrewGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
