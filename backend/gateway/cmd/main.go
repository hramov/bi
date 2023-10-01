package main

import (
	"context"
	"github.com/hramov/gvc-bi/backend/gateway/internal/api/http"
	"github.com/hramov/gvc-bi/backend/gateway/pkg/logger"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	l := logger.New("gateway", logger.Debug)

	httpServer := http.New(3010, l)
	httpServer.Start(ctx)
}
