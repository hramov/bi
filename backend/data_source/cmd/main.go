package main

import (
	"context"
	"github.com/hramov/gvc-bi/backend/datasource/internal/adapter/http"
	"github.com/hramov/gvc-bi/backend/datasource/pkg/database/postgres"
	"github.com/hramov/gvc-bi/backend/datasource/pkg/logger"
	"github.com/hramov/gvc-bi/backend/datasource/pkg/metrics"
	"github.com/joho/godotenv"
	"gopkg.in/Graylog2/go-gelf.v1/gelf"
	"io"
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

	writer := io.MultiWriter(os.Stderr)

	if os.Getenv("GRAYLOG_ADDR") != "" {
		graylogAddr := os.Getenv("GRAYLOG_ADDR")

		gelfWriter, err := gelf.NewWriter(graylogAddr)
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		writer = io.MultiWriter(os.Stderr, gelfWriter)
	}

	appLogger := logger.New("dashboard", logger.Debug, writer)

	metrics.InitMetrics()

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
