package plugin

import "context"

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

type Postgres interface {
	Connect(ctx context.Context, options ConnectOptions) (*string, error)
}
