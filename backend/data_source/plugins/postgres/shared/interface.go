package shared

import (
	"context"
	"github.com/hashicorp/go-plugin"
	postgres_go_plugin "github.com/hramov/postgres-datasource/proto"

	"google.golang.org/grpc"
)

var PluginName = "postgres"

type ConnectOptions struct {
	Dsn string
}

type ConnectResponse struct {
	Ok bool
	Id string
}

type CheckResponse struct {
	Ok bool
}

type QueryOptions struct {
	Sql string
}

type QueryResponse struct {
	Result string
}

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "POSTGRES_PLUGIN",
	MagicCookieValue: "postgres",
}

var PluginMap = map[string]plugin.Plugin{
	PluginName: &PostgresPlugin{},
}

type Postgres interface {
	Connect(ctx context.Context, options ConnectOptions) (*string, error)
}

type PostgresPlugin struct {
	plugin.Plugin
	Impl Postgres
}

func (p *PostgresPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	postgres_go_plugin.RegisterPostgresServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *PostgresPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (any, error) {
	return &GRPCClient{client: postgres_go_plugin.NewPostgresClient(c)}, nil
}
