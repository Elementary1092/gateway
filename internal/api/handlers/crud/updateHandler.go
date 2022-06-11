package crud

import (
	"context"
	"encoding/json"
	"github.com/elem1092/gateway/internal/domain"
	services "github.com/elem1092/gateway/pkg/client/grpc/CRUDService"
	"github.com/elem1092/gateway/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

// PATCH /posts/{id:[0-9]+} {"title": "...", "body": "..."}
func (u *updateHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json")

	encoder := json.NewEncoder(writer)

	vars := mux.Vars(request)
	idStr, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		u.logger.Errorf("Could not find id variable on: %s", request.URL.String())
		errMsg := &domain.ErrorDTO{Error: "no id"}
		if err := encoder.Encode(errMsg); err != nil {
			u.logger.Errorf("Error in encoding error message: %v", err)
		}
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		u.logger.Errorf("Could not convert id %s to int32", idStr)
		errMsg := &domain.ErrorDTO{Error: "invalid id"}
		if err := encoder.Encode(errMsg); err != nil {
			u.logger.Errorf("Error in encoding error message: %v", err)
		}
		return
	}

	content, err := u.parseBodyToContentDTO(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		errMsg := &domain.ErrorDTO{Error: "invalid request"}
		if err = encoder.Encode(errMsg); err != nil {
			u.logger.Errorf("Error in encoding error message: %v", err)
		}
		return
	}

	post := &services.UpdatePostDTO{
		Id:      int32(id),
		Content: content,
	}

	_, err = u.client.UpdatePost(context.Background(), post)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errMsg := &domain.ErrorDTO{Error: err.Error()}
		if err = encoder.Encode(errMsg); err != nil {
			u.logger.Errorf("Error in encoding error message: %v", err)
		}
		return
	}

	writer.WriteHeader(http.StatusOK)
	resp := map[string]string{"message": "success"}
	if err = encoder.Encode(resp); err != nil {
		u.logger.Errorf("Error while encoding response: %v", err)
	}
}

func (u *updateHandler) parseBodyToContentDTO(r *http.Request) (*services.ContentDTO, error) {
	decoder := json.NewDecoder(r.Body)
	post := &domain.UpdatePostDTO{}

	u.logger.Info("Decoding request body")
	if err := decoder.Decode(post); err != nil {
		u.logger.Errorf("Error in decoding request body: %v", err)
		return nil, err
	}

	return &services.ContentDTO{
		Title: post.Title,
		Body:  post.Body,
	}, nil
}
