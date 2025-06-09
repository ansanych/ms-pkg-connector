package connector

import (
	"context"
	"errors"
	"strconv"
	"sync"

	cerasus "github.com/ansanych/cerasus-proto/api_v2"
	config "github.com/ansanych/ms-pkg-config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientsPool struct {
	Pool map[string]*Client
}

type Client struct {
	Mu   sync.Mutex
	Conn *grpc.ClientConn
}

var Pool *ClientsPool

func InitPool(ctx context.Context, config *config.Config) error {
	Pool = new(ClientsPool)

	return Pool.BuildClients(ctx, config)
}

func (p *ClientsPool) BuildClients(ctx context.Context, config *config.Config) error {
	list, err := connectorClient.GetClientsList(ctx, &cerasus.Auth{})

	if err != nil {
		return err
	}

	clientsPool := make(map[string]*Client, len(config.Clients))

	for _, c := range config.Clients {

		find := false

		for _, l := range list.Data {
			if c == l.Service {
				conn, err := grpc.NewClient(l.Host+":"+strconv.FormatUint(uint64(l.Port), 10), grpc.WithTransportCredentials(insecure.NewCredentials()))

				if err != nil {
					return err
				}

				client := &Client{
					Conn: conn,
				}

				clientsPool[c] = client

				find = true
			}
		}

		if !find {
			return errors.New(c + "not found client in connector")
		}
	}

	p.Pool = clientsPool

	return nil
}

func (p *ClientsPool) GetClient(service string) (*Client, error) {

	client, ok := p.Pool[service]

	if !ok {
		return nil, errors.New("not found client")
	}

	client.Mu.Lock()

	return client, nil
}
