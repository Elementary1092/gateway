package fetcher

import (
	"context"
	"encoding/json"
	"github.com/elem1092/gateway/internal/domain"
	fetch "github.com/elem1092/gateway/pkg/client/grpc/FetcherService"
	"github.com/elem1092/gateway/pkg/logging"
	"net/http"
	"time"
)

type statusHandler struct {
	logger *logging.Logger
	client fetch.FetchServiceClient
}

func NewStatusHandler(logger *logging.Logger, client fetch.FetchServiceClient) http.Handler {
	return &statusHandler{
		logger: logger,
		client: client,
	}
}

// GET /status
func (s *statusHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text-json")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	out, err := s.client.GetStatus(ctx, &fetch.EmptyMessage{})
	encoder := json.NewEncoder(writer)
	if err != nil {
		errMsg := domain.ErrorDTO{Error: err.Error()}
		if err := encoder.Encode(errMsg); err != nil {
			s.logger.Errorf("Got an error from encoder: %v", err)
		}
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := encoder.Encode(domain.StatusDTO{
		Status: fetch.Status_name[int32(out.GetStatusCode())],
	}); err != nil {
		s.logger.Errorf("Got an error from encoder: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
