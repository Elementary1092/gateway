package fetcher

import (
	"context"
	"encoding/json"
	"github.com/elem1092/gateway/internal/domain"
	fetch "github.com/elem1092/gateway/pkg/client/grpc/FetcherService"
	"github.com/elem1092/gateway/pkg/logging"
	"net/http"
)

type startFetchHandler struct {
	logger *logging.Logger
	client fetch.FetchServiceClient
}

func NewStartFetchHandler(logger *logging.Logger, client fetch.FetchServiceClient) http.Handler {
	return &startFetchHandler{
		logger: logger,
		client: client,
	}
}

// GET /
func (s *startFetchHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json")

	s.logger.Info("Starting fetching process")

	encoder := json.NewEncoder(writer)
	status, err := s.client.StartFetching(context.Background(), &fetch.FetchRequest{})
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		s.logger.Errorf("Got error while start of fetch process: %v", err)
		errMsg := &domain.ErrorDTO{Error: err.Error()}
		if err = encoder.Encode(errMsg); err != nil {
			s.logger.Errorf("Got error while encoding response: %v", err)
		}
		return
	}

	if err = encoder.Encode(status); err != nil {
		s.logger.Errorf("Got error while encodeing response: %v", err)
	}
	writer.WriteHeader(http.StatusOK)
}
