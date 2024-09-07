package main

import (
	"github.com/peygy/medapp/internal/pkg/context"
	"github.com/peygy/medapp/internal/pkg/gin"
	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/peygy/medapp/internal/services/graphql/config"
	"github.com/peygy/medapp/internal/services/graphql/internal/configurations"
	"github.com/peygy/medapp/internal/services/graphql/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewGraphQLConfig,
				logger.NewLogger,
				context.NewContext,
				gin.NewGinServer,
				grpc.NewGrpcClient,
			),
			fx.Invoke(configurations.InitEndpoints),
			fx.Invoke(server.RunServers),
		),
	).Run()
}

/*
go get github.com/99designs/gqlgen/codegen/config@v0.17.49
go get github.com/99designs/gqlgen/internal/imports@v0.17.49
go get github.com/99designs/gqlgen@v0.17.49
go run github.com/99designs/gqlgen generate
*/
