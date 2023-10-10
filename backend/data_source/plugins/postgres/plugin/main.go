package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/go-plugin"
	"github.com/hramov/postgres-datasource/shared"
	_ "github.com/lib/pq"
)

type Postgres struct {
	conns map[string]*sql.DB
}

func (p *Postgres) Connect(ctx context.Context, opts shared.ConnectOptions) (*string, error) {

	if p.conns == nil {
		p.conns = make(map[string]*sql.DB)
	}

	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to postgres (%s): %v", opts.Dsn, err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot ping postgres (%s): %v", opts.Dsn, err)
	}
	id := uuid.New()
	idPtr := id.String()
	p.conns[idPtr] = db
	return &idPtr, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			shared.PluginName: &shared.PostgresPlugin{Impl: &Postgres{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
