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

type fetcherErrorHandler struct {
	logger *logging.Logger
	client fetch.FetchServiceClient
}

func NewFetcherHandler(logger *logging.Logger, client fetch.FetchServiceClient) http.Handler {
	return &fetcherErrorHandler{
		logger: logger,
		client: client,
	}
}

// GET /error
func (f *fetcherErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	f.logger.Info("Handling /error request")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := f.client.GetError(ctx, &fetch.EmptyMessage{})
	errMsg := domain.ErrorDTO{Error: "No error"}

	encoder := json.NewEncoder(writer)
	if err != nil {
		f.logger.Infof("Error occurred in Fetcher service: %v", err)
		errMsg.Error = err.Error()
	}

	if err = encoder.Encode(errMsg); err != nil {
		f.logger.Errorf("Error occurred while encoding json: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
