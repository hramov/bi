package shared

import (
	"context"
	postgres_go_plugin "github.com/hramov/postgres-datasource/proto"
)

type GRPCClient struct {
	client postgres_go_plugin.PostgresClient
}

func (m *GRPCClient) Connect(ctx context.Context, opts ConnectOptions) (*string, error) {
	resp, err := m.client.Connect(ctx, &postgres_go_plugin.ConnectRequest{
		Dsn: opts.Dsn,
	})

	if err != nil {
		return nil, err
	}

	id := resp.Id

	return &id, nil
}

type GRPCServer struct {
	postgres_go_plugin.UnimplementedPostgresServer
	Impl Postgres
}

func (m *GRPCServer) Connect(ctx context.Context, req *postgres_go_plugin.ConnectRequest) (*postgres_go_plugin.ConnectResponse, error) {
	id, err := m.Impl.Connect(ctx, ConnectOptions{
		Dsn: req.Dsn,
	})

	if err != nil {
		return &postgres_go_plugin.ConnectResponse{
			Ok: false,
		}, err
	}

	return &postgres_go_plugin.ConnectResponse{
		Ok: true,
		Id: *id,
	}, err
}
