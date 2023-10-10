package plugin

import (
	"context"
	"net/rpc"
)

type PluginServerRPC struct {
	Impl Postgres
}

func (s *PluginServerRPC) Connect(ctx context.Context, opts ConnectOptions) (*string, error) {
	return s.Impl.Connect(ctx, opts)
}

type PluginClientRPC struct {
	client *rpc.Client
}

func (c *PluginClientRPC) Connect(ctx context.Context, opts ConnectOptions) (*string, error) {
	var id = new(string)
	if err := c.client.Call("Plugin.Connect", opts, id); err != nil {
		return nil, err
	}
	return id, nil
}
