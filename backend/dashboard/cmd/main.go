package main

import (
	"context"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/api/http"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/api/http_ds"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/database/postgres"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/logger"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	l := logger.New("dashboard", logger.Debug)

	pg, err := postgres.New(&postgres.Options{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "gvc_bi",
		SslMode:  "disable",
	}, "")

	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}

	s := http.New(3010, pg, l)
	go s.Start(ctx)

	dsServer := http_ds.New(3011, pg, l)
	err = dsServer.Start(ctx)
	if err != nil {
		l.Error(err.Error())
		os.Exit(0)
	}
}
