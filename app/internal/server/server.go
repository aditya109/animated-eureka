// Package sample http API
//
// Documentation of crud-template API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta

package server

import (
	"fmt"
	"net/http"

	"github.com/aditya109/animated-eureka/internal/router"
	logger "github.com/sirupsen/logrus"
)

func Start() {
	// configuring router for the server
	router := router.ConfigureRouter()
	logger.Info("router configuration successful")
	logger.Info(fmt.Sprintf("starting server at %s://%s", prefix, endpoint))
	logger.Info(fmt.Sprintf("swagger docs can be viewed at %s://%s/docs", prefix, endpoint))

	// configuring server
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", httpPort),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	logger.Fatal(srv.ListenAndServe())
}
