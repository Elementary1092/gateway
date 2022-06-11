package handlers

import (
	"context"
	"github.com/elem1092/gateway/internal/api/handlers/crud"
	"github.com/elem1092/gateway/internal/api/handlers/fetcher"
	"github.com/elem1092/gateway/internal/config"
	fetch "github.com/elem1092/gateway/pkg/client/grpc/FetcherService"
	crudClient "github.com/elem1092/gateway/pkg/client/grpc/crud"
	fetcherClient "github.com/elem1092/gateway/pkg/client/grpc/fetcher"
	"github.com/elem1092/gateway/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeRouter(logger *logging.Logger, cfg *config.Configuration) http.Handler {
	logger.Info("Starting router initialization")

	crudClient := crudClient.NewCrudServiceClientWrapper(logger, cfg.CRUDCfg)
	fetchClient := fetcherClient.NewCrudServiceClientWrapper(logger, cfg.FetcherCfg)

	getStatusHandler := fetcher.NewStatusHandler(logger, fetchClient)
	getErrorHandler := fetcher.NewFetcherHandler(logger, fetchClient)
	getAllHandler := crud.NewGetAllHandler(logger, crudClient)
	getPostHandler := crud.NewGetPostHandler(logger, crudClient)
	getUserPostsHandler := crud.NewGetUserPostsHandler(logger, crudClient)
	deletePostHandler := crud.NewDeleteHandler(logger, crudClient)
	savePostHandler := crud.NewSaveHandler(logger, crudClient)
	updatePostHandler := crud.NewUpdateHandler(logger, crudClient)

	router := mux.NewRouter()

	router.Handle("/status", getStatusHandler).Methods("GET")
	router.Handle("/error", getErrorHandler).Methods("GET")
	router.Handle("/posts", getAllHandler).Methods("GET")
	router.Handle("/posts/{id:[0-9]+}", getPostHandler).Methods("GET")
	router.Handle("/posts/user/{id:[0-9]+}", getUserPostsHandler).Methods("GET")
	router.Handle("/posts/{id:[0-9]+}", deletePostHandler).Methods("DELETE")
	router.Handle("/posts", savePostHandler).Methods("POST")
	router.Handle("/posts/{id:[0-9]+}", updatePostHandler).Methods("PATCH")

	logger.Info("Finished router initialization")

	logger.Info("Starting fetch process")
	fetchClient.StartFetching(context.Background(), &fetch.FetchRequest{})

	return router
}
