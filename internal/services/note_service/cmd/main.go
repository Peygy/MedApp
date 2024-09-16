package main

import (
	"go.uber.org/fx"

	"github.com/peygy/medapp/internal/pkg/context"
	"github.com/peygy/medapp/internal/pkg/database/postgres"
	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"

	"github.com/peygy/medapp/internal/services/note_service/config"
	"github.com/peygy/medapp/internal/services/note_service/internal/data"
	grpcConn "github.com/peygy/medapp/internal/services/note_service/internal/grpc"
	"github.com/peygy/medapp/internal/services/note_service/internal/services"
	"github.com/peygy/medapp/internal/services/note_service/server"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewNoteConfig,
				logger.NewLogger,
				context.NewContext,
				grpc.NewGrpcServer,
				postgres.NewDatabaseConnection,

				services.NewHealthService,
			),
			fx.Invoke(data.InitDatabaseSchema),
			fx.Invoke(grpcConn.InitNoteGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
