package factory

import (
	"fmt"
	"github.com/TemaKut/messenger-auth/internal/app/config"
	usergrpchandler "github.com/TemaKut/messenger-auth/internal/app/handler/grpc/user"
	"github.com/TemaKut/messenger-auth/internal/app/logger"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"net"
	"time"
)

var GrpcSet = wire.NewSet(
	ProvideGrpcProvider,
	ProvideGrpcServer,
)

type GrpcProvider struct{}

func ProvideGrpcProvider(
	_ GrpcServerProvider,
) GrpcProvider {
	return GrpcProvider{}
}

type GrpcServerProvider struct{}

func ProvideGrpcServer(
	cfg *config.Config,
	userGrpcHandler *usergrpchandler.Handler,
	logger *logger.Logger,
) (GrpcServerProvider, func(), error) {
	logger.Infof("start grpc server on %s", cfg.Server.Grpc.Addr)

	lis, err := net.Listen("tcp", cfg.Server.Grpc.Addr)
	if err != nil {
		return GrpcServerProvider{}, nil, fmt.Errorf("error listen addr %s. %w", cfg.Server.Grpc.Addr, err)
	}

	server := grpc.NewServer()

	authv1.RegisterUserAPIServer(server, userGrpcHandler)

	errCh := make(chan error, 1)

	go func() {
		if err := server.Serve(lis); err != nil {
			errCh <- fmt.Errorf("error serve grpc server %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return GrpcServerProvider{}, nil, fmt.Errorf("error from errCh %w", err)
	case <-time.After(200 * time.Millisecond):
	}

	return GrpcServerProvider{}, func() {
		logger.Infof("shutdown grpc server on %s", cfg.Server.Grpc.Addr)

		server.GracefulStop()
	}, nil
}
