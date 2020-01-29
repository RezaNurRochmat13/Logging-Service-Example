package main

import (
	"log"
	"net/http"
	"svc-logger-go/config"
	// Log Requests Packages
	LogRequestsHandlerPackage "svc-logger-go/modules/v1/log-request/presenter"
	LogRequestsRepoPackage "svc-logger-go/modules/v1/log-request/repository"
	LogRequestsUseCasePackage "svc-logger-go/modules/v1/log-request/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	mongoConnection, errorMongoConn := config.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}

	echoRouter := echo.New()
	echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Import module log request
	logRequestsRepo := LogRequestsRepoPackage.NewLogRequestRepository(mongoConnection)
	logRequestsUseCase := LogRequestsUseCasePackage.NewLogRequestUseCase(logRequestsRepo)
	LogRequestsHandlerPackage.NewLogRequestHandler(echoRouter, logRequestsUseCase)

	//Configuration of logger
	echoRouter.Use(middleware.Logger())
	echoRouter.Logger.Fatal(echoRouter.Start(":8081"))
}
