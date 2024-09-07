package main

import (
	"github.com/peygy/medapp/internal/pkg/context"
	"github.com/peygy/medapp/internal/pkg/database/postgres"
	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/peygy/medapp/internal/services/auth_service/config"
	"github.com/peygy/medapp/internal/services/auth_service/internal/data"
	grpcConn "github.com/peygy/medapp/internal/services/auth_service/internal/grpc"
	"github.com/peygy/medapp/internal/services/auth_service/internal/managers"
	"github.com/peygy/medapp/internal/services/auth_service/server"
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

				managers.NewRoleManager,
				managers.NewUserManager,
			),
			fx.Invoke(data.InitDatabaseSchema),
			fx.Invoke(grpcConn.InitAuthGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
