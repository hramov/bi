package main

import (
	"context"
	"fmt"
	"github.com/hramov/gvc-bi/backend/datasource/plugins/postgres-rpc/plugin"
	"os"
)

func main() {
	// Load plugins from the plugin-binaries directory.
	var pm plugin.Manager
	pl, err := pm.LoadPlugins("./bin/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer pm.Close()

	id, err := pl.Connect(context.Background(), plugin.ConnectOptions{Dsn: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"})

	fmt.Println(id)

}
