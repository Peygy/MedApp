package server

import (
	"context"

	"github.com/peygy/medapp/internal/pkg/gin"
	"github.com/peygy/medapp/internal/pkg/logger"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, ctx context.Context, log logger.ILogger, gin *gin.GinServer) error {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := gin.Run(ctx); err != nil {
					log.Fatal("Error running gin server")
				}
			}()

			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Info("All servers shutdown gracefully...")
			return nil
		},
	})

	return nil
}
