package connector

import (
	"context"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cerasus "github.com/ansanych/cerasus-proto/api_v2"
	config "github.com/ansanych/ms-pkg-config"
)

var connectorClient cerasus.ConnectorClient

func RunConnectorClient(ctx context.Context, config *config.Config) error {
	conn, err := grpc.NewClient(config.Connector.Host+":"+strconv.Itoa(config.Connector.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	connectorClient = cerasus.NewConnectorClient(conn)

	return nil
}
