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

var (
    ErrNoUserIdField = errors.New("could not find user_id field in request body")
)

type getHandler struct {
    logger *logging.Logger
    client services.CRUDServiceClient
}

func NewGetHandler(logger *logging.Logger, client services.CRUDServiceClient) http.Handler {
    return &getHandler{
        logger: logger,
        client: client,
    }
}

// GET /posts?id=   {id_type: '...'} id_type may be 'USER_ID', 'POST_ID'
func (g *getHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "text/json")
    g.logger.Info("Processing GET ALL POSTS BY ID request")

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
        Needed: services.GetPostsRequest_POST_ID,
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

    post, err := client.Recv()
    if err != nil {
        if errors.Is(err, io.EOF) {
            g.logger.Errorf("No post with id: %d", id)
            return
        }

        g.logger.Errorf("Got error while parsing posts: %v", err)
        return
    }

    if err = encoder.Encode(post); err != nil {
        g.logger.Errorf("Got error while encoding post: %v", err)
        writer.WriteHeader(http.StatusInternalServerError)
        return
    }

    writer.WriteHeader(http.StatusOK)
}
