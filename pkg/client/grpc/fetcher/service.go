package fetcher

import (
	"context"
	"fmt"
	"github.com/elem1092/gateway/internal/config"
	fetch "github.com/elem1092/gateway/pkg/client/grpc/FetcherService"
	"github.com/elem1092/gateway/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fetcherServiceClient struct {
	logger  *logging.Logger
	service fetch.FetchServiceClient
}

func NewCrudServiceClientWrapper(logger *logging.Logger, cfg config.FetcherServiceConfig) fetch.FetchServiceClient {
	logger.Infof("Connecting to a service on %s:%s", cfg.Address, cfg.Port)
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Errorf("Unable to connect to a service due to: %v", err)
		return nil
	}

	logger.Infof("Connected to a service on %s:%s successfully", cfg.Address, cfg.Port)
	service := fetch.NewFetchServiceClient(conn)

	return &fetcherServiceClient{logger, service}
}

func (f *fetcherServiceClient) StartFetching(ctx context.Context,
	in *fetch.FetchRequest,
	opts ...grpc.CallOption) (*fetch.FetchStatus, error) {
	f.logger.Info("Handling StartFetching request")

	out, err := f.service.StartFetching(ctx, in, opts...)
	if err != nil {
		f.logger.Errorf("Caught an error from a server: %v", err)
		return nil, err
	}

	return out, nil
}

func (f *fetcherServiceClient) GetStatus(ctx context.Context,
	in *fetch.EmptyMessage,
	opts ...grpc.CallOption) (*fetch.FetchStatus, error) {
	f.logger.Info("Handling GetStatus request")

	out, err := f.service.GetStatus(ctx, in, opts...)
	if err != nil {
		f.logger.Errorf("Caught an error from a server: %v", err)
		return nil, err
	}

	return out, nil
}

func (f *fetcherServiceClient) GetError(ctx context.Context,
	in *fetch.EmptyMessage,
	opts ...grpc.CallOption) (*fetch.EmptyMessage, error) {
	f.logger.Info("Handling GetError request")

	out, err := f.service.GetError(ctx, in, opts...)
	if err != nil {
		f.logger.Errorf("Caught an error from a server: %v", err)
		return nil, err
	}

	return out, nil
}
