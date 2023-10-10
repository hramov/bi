package main

import (
	"context"
	"fmt"
	"github.com/hramov/postgres-datasource/shared"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

func main() {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		Plugins:          shared.PluginMap,
		Cmd:              exec.Command("sh", "-c", "./bin/postgres"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	err = rpcClient.Ping()
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	raw, err := rpcClient.Dispense(shared.PluginName)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	pg, ok := raw.(shared.Postgres)
	if !ok {
		fmt.Println("cannot cast interface")
		os.Exit(1)
	}

	result, err := pg.Connect(context.Background(), shared.ConnectOptions{
		Dsn: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(*result)

	os.Exit(0)
}
