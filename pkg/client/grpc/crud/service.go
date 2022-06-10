package crud

import (
	"context"
	"fmt"
	"github.com/elem1092/gateway/internal/config"
	services "github.com/elem1092/gateway/pkg/client/grpc/CRUDService"
	"github.com/elem1092/gateway/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type crudServiceClient struct {
	logger  *logging.Logger
	service services.CRUDServiceClient
}

func NewCrudServiceClientWrapper(logger *logging.Logger, cfg config.CRUDServiceConfig) services.CRUDServiceClient {
	logger.Infof("Connecting to a service on %s:%s", cfg.Address, cfg.Port)
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Errorf("Unable to connect to a service due to: %v", err)
		return nil
	}

	logger.Infof("Connected to a service on %s:%s successfully", cfg.Address, cfg.Port)
	service := services.NewCRUDServiceClient(conn)

	return &crudServiceClient{logger, service}
}

func (c *crudServiceClient) SavePost(ctx context.Context,
	in *services.SavePostDTO,
	opts ...grpc.CallOption) (*services.PostDTO, error) {
	c.logger.Info("Handling SavePost request")

	out, err := c.service.SavePost(ctx, in, opts...)
	if err != nil {
		c.logger.Errorf("Caught an error from a server: %v", err)
		return nil, err
	}

	return out, nil
}

func (c *crudServiceClient) GetPosts(ctx context.Context,
	in *services.GetPostsRequest,
	opts ...grpc.CallOption) (services.CRUDService_GetPostsClient, error) {
	c.logger.Info("Handling GetPosts request")

	out, err := c.service.GetPosts(ctx, in, opts...)
	if err != nil {
		c.logger.Errorf("Caught an error from a server: %v", err)
		return nil, err
	}

	return out, nil
}

func (c *crudServiceClient) DeletePost(ctx context.Context,
	in *services.DeleteRequest,
	opts ...grpc.CallOption) (*services.ErrorResponse, error) {
	c.logger.Info("Handling DeletePost request")

	out, err := c.service.DeletePost(ctx, in, opts...)
	if err != nil {
		c.logger.Errorf("Caught an error from a server: %v", err)
		return nil, err
	}

	return out, nil
}

func (c *crudServiceClient) UpdatePost(ctx context.Context,
	in *services.UpdatePostDTO,
	opts ...grpc.CallOption) (*services.ErrorResponse, error) {
	c.logger.Info("Handling UpdatePost request")

	out, err := c.service.UpdatePost(ctx, in, opts...)
	if err != nil {
		c.logger.Errorf("Caught an error from a server: %v", err)
		return nil, err
	}

	return out, nil
}
