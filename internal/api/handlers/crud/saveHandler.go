package crud

import (
	services "github.com/elem1092/gateway/pkg/client/grpc/CRUDService"
	"github.com/elem1092/gateway/pkg/logging"
	"net/http"
)

type saveHandler struct {
	logger *logging.Logger
	client services.CRUDServiceClient
}

func NewSaveHandler(logger *logging.Logger, client services.CRUDServiceClient) http.Handler {
	return &saveHandler{
		logger: logger,
		client: client,
	}
}

// POST /posts {user_id: ..., title: '...', body: '...'}
func (s *saveHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json")
	//TODO implement me
	panic("implement me")
}

func (s *saveHandler) parseBodyToSaveDTO(r *http.Request) (*services.SavePostDTO, error) {

	return nil, nil
}
