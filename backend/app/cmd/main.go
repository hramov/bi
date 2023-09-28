package main

import (
	"context"
	"github.com/hramov/gvc-bi/backend/internal/api/http"
	"github.com/hramov/gvc-bi/backend/pkg/database/postgres"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	pg, err := postgres.New(postgres.Options{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "gvc_bi",
		SslMode:  "disable",
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	s := http.New(3010, pg)
	s.Start(ctx)
}
