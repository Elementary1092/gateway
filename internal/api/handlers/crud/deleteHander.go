package crud

import (
    "context"
    "encoding/json"
    "github.com/elem1092/gateway/internal/domain"
    services "github.com/elem1092/gateway/pkg/client/grpc/CRUDService"
    "github.com/elem1092/gateway/pkg/logging"
    "net/http"
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

// DELETE /posts/id=? {id: ...}
func (d *deleteHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    d.logger.Info("Handling delete request")

    encoder := json.NewEncoder(writer)
    reader, err := request.GetBody()
    if err != nil {
        d.logger.Errorf("Could not get body of a delete request due to: %v", err)
        errMsg := &domain.ErrorDTO{Error: "Invalid request"}
        writer.WriteHeader(http.StatusBadRequest)
        if err = encoder.Encode(errMsg); err != nil {
            d.logger.Errorf("could not write response due to: %v", err)
        }
        return
    }

    d.logger.Info("Decoding request")

    dto := domain.DeleteDTO{}
    decoder := json.NewDecoder(reader)
    if err = decoder.Decode(dto); err != nil {
        d.logger.Errorf("Failed to decode delete request due to: %v", err)
        writer.WriteHeader(http.StatusBadRequest)
        return
    }

    req := &services.DeleteRequest{Id: dto.Id}
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
}
