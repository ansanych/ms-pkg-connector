package connector

import (
	"context"

	cerasus "github.com/ansanych/cerasus-proto/api_v2"
	config "github.com/ansanych/ms-pkg-config"
)

func RegisterServerOnConnector(ctx context.Context, config *config.Config) error {

	_, err := connectorClient.SetClientAddress(ctx, &cerasus.Client{
		Service: config.Service,
		Host:    config.Address.Host,
		Port:    uint32(config.Address.Port),
	})

	return err
}
