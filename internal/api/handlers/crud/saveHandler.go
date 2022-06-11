package crud

import (
	"context"
	"encoding/json"
	"github.com/elem1092/gateway/internal/domain"
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

// POST /posts {"user_id": ..., "title": "...", "body": "..."}
func (s *saveHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json")

	encoder := json.NewEncoder(writer)

	newPost, err := s.parseBodyToSaveDTO(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		errMsg := domain.ErrorDTO{Error: err.Error()}
		if err = encoder.Encode(errMsg); err != nil {
			s.logger.Errorf("Error while encoding json: %v", err)
		}
		return
	}

	savedPost, err := s.client.SavePost(context.Background(), newPost)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errMsg := domain.ErrorDTO{Error: err.Error()}
		if err = encoder.Encode(errMsg); err != nil {
			s.logger.Errorf("Error while encoding json: %v", err)
		}
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := encoder.Encode(savedPost); err != nil {
		s.logger.Errorf("Error while encoding response: %v", err)
	}
}

func (s *saveHandler) parseBodyToSaveDTO(r *http.Request) (*services.SavePostDTO, error) {
	s.logger.Info("Parsing request body")
	decoder := json.NewDecoder(r.Body)

	newPost := &domain.SavePostDTO{}
	if err := decoder.Decode(newPost); err != nil {
		s.logger.Errorf("error while decoding request body: %v", err)
		return nil, err
	}

	return &services.SavePostDTO{
		UserId: newPost.UserId,
		Content: &services.ContentDTO{
			Title: newPost.Title,
			Body:  newPost.Body,
		},
	}, nil
}
