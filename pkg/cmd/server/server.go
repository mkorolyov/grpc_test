package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/mkorolyov/grpc_test/pkg/protocol/grpc"
	"github.com/mkorolyov/grpc_test/pkg/service/v1"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	GRPCPort := flag.String("grpc-port", "", "gRPC port to bind")
	flag.Parse()

	if len(*GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", *GRPCPort)
	}

	v1API := v1.NewToDoServiceServer()

	return grpc.RunServer(ctx, v1API, *GRPCPort)
}
