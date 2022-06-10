package crud

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elem1092/gateway/internal/domain"
	services "github.com/elem1092/gateway/pkg/client/grpc/CRUDService"
	"github.com/elem1092/gateway/pkg/logging"
	"io"
	"net/http"
)

type getAllHandler struct {
	logger *logging.Logger
	client services.CRUDServiceClient
}

func NewGetAllHandler(logger *logging.Logger, client services.CRUDServiceClient) http.Handler {
	return &getAllHandler{
		logger: logger,
		client: client,
	}
}

// GET /posts
func (g *getAllHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json")
	g.logger.Info("Processing GET ALL POSTS request")
	req := &services.GetPostsRequest{
		Id:     0,
		Needed: services.GetPostsRequest_ALL,
	}

	encoder := json.NewEncoder(writer)
	client, err := g.client.GetPosts(context.Background(), req)
	if err != nil {
		g.logger.Errorf("Got error while getting all messages: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		errMsg := domain.ErrorDTO{Error: err.Error()}
		if err = encoder.Encode(errMsg); err != nil {
			g.logger.Errorf("Error while encoding error message: %v", err)
		}
		return
	}

	for {
		post, err := client.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			g.logger.Errorf("Got error while parsing posts: %v", err)
			return
		}

		if err = encoder.Encode(post); err != nil {
			g.logger.Errorf("Got error while encoding post: %v", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	writer.WriteHeader(http.StatusOK)
}
