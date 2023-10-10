package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type PostgresPlugin struct {
	Impl Postgres
}

func (p *PostgresPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &PluginServerRPC{
		Impl: p.Impl,
	}, nil
}

func (p *PostgresPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginClientRPC{
		client: c,
	}, nil
}
