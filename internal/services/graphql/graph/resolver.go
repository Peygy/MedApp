package graph

import (
	"github.com/peygy/medapp/internal/pkg/grpc"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GrpcServices []grpc.GrpcService
}
