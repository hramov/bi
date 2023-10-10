package plugin

import (
	"fmt"
	goplugin "github.com/hashicorp/go-plugin"
	"os/exec"
)

type Manager struct {
	pluginClients []*goplugin.Client
}

// LoadPlugins takes a directory path and assumes that all files within it
// are plugin binaries. It runs all these binaries in sub-processes,
// establishes RPC communication with the plugins, and registers them for
// the hooks they declare to support.
func (m *Manager) LoadPlugins(path string) (Postgres, error) {
	binaries, err := goplugin.Discover("*", path)
	if err != nil {
		return nil, err
	}

	pluginMap := map[string]goplugin.Plugin{
		PluginName: &PostgresPlugin{},
	}

	for _, bpath := range binaries {
		client := goplugin.NewClient(&goplugin.ClientConfig{
			HandshakeConfig: Handshake,
			Plugins:         pluginMap,
			Cmd:             exec.Command(bpath),
		})
		m.pluginClients = append(m.pluginClients, client)

		rpcClient, err := client.Client()
		if err != nil {
			return nil, err
		}

		raw, err := rpcClient.Dispense(PluginName)
		if err != nil {
			return nil, err
		}

		impl := raw.(Postgres)

		return impl, nil
	}

	return nil, fmt.Errorf("cannot find binary")
}

func (m *Manager) Close() {
	for _, client := range m.pluginClients {
		client.Kill()
	}
}
