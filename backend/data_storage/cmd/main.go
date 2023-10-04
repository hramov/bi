package main

import (
	"context"
	"github.com/hramov/gvc-bi/backend/datastorage/internal/adapter/http"
	"github.com/hramov/gvc-bi/backend/datastorage/pkg/database/postgres"
	"github.com/hramov/gvc-bi/backend/datastorage/pkg/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strconv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	appLogger := logger.New("dashboard", logger.Debug)

	pg, err := postgres.New(nil, os.Getenv("PG_DSN"))

	if err != nil {
		appLogger.Error(err.Error())
		os.Exit(1)
	}

	portStr := os.Getenv("DATA_SOURCE_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		appLogger.Error(err.Error())
		os.Exit(1)
	}

	dsServer := http.New(port, pg, appLogger)
	err = dsServer.Start(ctx)
	if err != nil {
		appLogger.Error(err.Error())
		os.Exit(0)
	}
}
