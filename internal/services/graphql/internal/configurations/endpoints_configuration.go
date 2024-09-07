package configurations

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginServer "github.com/peygy/medapp/internal/pkg/gin"
	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/services/graphql/graph"
	"github.com/peygy/medapp/internal/services/graphql/internal/middleware"
)

func InitEndpoints(eng *ginServer.GinServer, grpcPull *grpc.GrpcPull) {
	eng.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	  }))

	eng.Engine.POST("/graphql/signup", graphqlHandler(grpcPull.Services))
	eng.Engine.POST("/graphql/signin", graphqlHandler(grpcPull.Services))

	protected := eng.Engine.Group("/graphql")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/account", graphqlHandler(grpcPull.Services))
		protected.POST("/logout", graphqlHandler(grpcPull.Services))
	}
}

// Defining the Graphql handler
func graphqlHandler(grpcServices []grpc.GrpcService) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{GrpcServices: grpcServices}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
