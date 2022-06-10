package crud

import (
	services "github.com/elem1092/gateway/pkg/client/grpc/CRUDService"
	"github.com/elem1092/gateway/pkg/logging"
	"net/http"
)

type updateHandler struct {
	logger *logging.Logger
	client services.CRUDServiceClient
}

func NewUpdateHandler(logger *logging.Logger, client services.CRUDServiceClient) http.Handler {
	return &updateHandler{
		logger: logger,
		client: client,
	}
}

// PATCH /posts/id=? {title: '...', body: '...'}
func (u updateHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
