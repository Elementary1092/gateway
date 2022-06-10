package main

import (
    "github.com/elem1092/gateway/internal/api/handlers/crud"
    "github.com/elem1092/gateway/internal/api/handlers/fetcher"
    "github.com/elem1092/gateway/internal/config"
    crudClient "github.com/elem1092/gateway/pkg/client/grpc/crud"
    fetcherClient "github.com/elem1092/gateway/pkg/client/grpc/fetcher"
    "github.com/elem1092/gateway/pkg/logging"
    "github.com/gorilla/mux"
)

func main() {
    logger := logging.GetLogger()
    logger.Info("Starting API gateway")

    cfg := config.GetConfiguration()

    crudClient := crudClient.NewCrudServiceClientWrapper(logger, cfg.CRUDCfg)
    fetchClient := fetcherClient.NewCrudServiceClientWrapper(logger, cfg.FetcherCfg)

    getStatusHandler := fetcher.NewStatusHandler(logger, fetchClient)
    getErrorHandler := fetcher.NewFetcherHandler(logger, fetchClient)
    getAllHandler := crud.NewGetAllHandler(logger, crudClient)

    router := mux.NewRouter()

    router.Handle("/status", getStatusHandler)
    router.Handle("/error", getErrorHandler)
    router.Handle("/posts", getAllHandler)
}
