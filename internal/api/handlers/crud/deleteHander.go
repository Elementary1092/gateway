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

type deleteHandler struct {
	logger *logging.Logger
	client services.CRUDServiceClient
}

func NewDeleteHandler(logger *logging.Logger, client services.CRUDServiceClient) http.Handler {
	return &deleteHandler{
		logger: logger,
		client: client,
	}
}

// DELETE /posts?id=
func (d *deleteHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json")
	d.logger.Info("Handling delete request")

	encoder := json.NewEncoder(writer)

	vars := mux.Vars(request)
	idStr, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		errMsg := &domain.ErrorDTO{Error: "Invalid request"}
		writer.WriteHeader(http.StatusBadRequest)
		if err := encoder.Encode(errMsg); err != nil {
			d.logger.Errorf("could not write response due to: %v", err)
		}
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		errMsg := &domain.ErrorDTO{Error: "Invalid request; expected id to be integer type"}
		writer.WriteHeader(http.StatusBadRequest)
		if err := encoder.Encode(errMsg); err != nil {
			d.logger.Errorf("could not write response due to: %v", err)
		}
		return
	}

	req := &services.DeleteRequest{Id: int32(id)}
	_, err = d.client.DeletePost(context.Background(), req)
	if err != nil {
		d.logger.Errorf("Could not delete record due to: %v", err)
		errMsg := &domain.ErrorDTO{Error: err.Error()}
		writer.WriteHeader(http.StatusInternalServerError)
		if err = encoder.Encode(errMsg); err != nil {
			d.logger.Errorf("could not write response due to: %v", err)
		}
		return
	}

	writer.WriteHeader(http.StatusOK)
	resp := map[string]string{"message": "success"}
	if err = encoder.Encode(resp); err != nil {
		d.logger.Errorf("Error while encoding response: %v", err)
	}
}
