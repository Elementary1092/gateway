package crud

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elem1092/gateway/internal/domain"
	services "github.com/elem1092/gateway/pkg/client/grpc/CRUDService"
	"github.com/elem1092/gateway/pkg/logging"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type getUserPostsHandler struct {
	logger *logging.Logger
	client services.CRUDServiceClient
}

func NewGetUserPostsHandler(logger *logging.Logger, client services.CRUDServiceClient) http.Handler {
	return &getUserPostsHandler{
		logger: logger,
		client: client,
	}
}

// GET /posts/user/{id:[0-9}+}
func (g *getUserPostsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json")
	g.logger.Info("Processing GET ALL POSTS BY USER ID request")

	encoder := json.NewEncoder(writer)

	vars := mux.Vars(request)
	idStr, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		g.logger.Errorf("Could not find id variable on: %s", request.URL.String())
		errMsg := &domain.ErrorDTO{Error: "no id variable"}
		if err := encoder.Encode(errMsg); err != nil {
			g.logger.Errorf("Error in encoding error message: %v", err)
		}
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 32)

	req := &services.GetPostsRequest{
		Id:     int32(id),
		Needed: services.GetPostsRequest_USER_ID,
	}
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
