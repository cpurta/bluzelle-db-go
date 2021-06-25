package bluzelledbgo

import (
	"context"
	"net"

	tmnet "github.com/tendermint/tendermint/libs/net"
	"google.golang.org/grpc"
)

//go:generate mockgen -source ./client.go -destination ./mock_client/mock.go

// BluezelleClient provides a high level interface to access the Query and Transaction
// rpc methods via Querier and Transactioner interfaces.
type BluezelleClient interface {
	Query() QueryClient
	Transaction() TransactionClient
}

type Config struct {
	Mnemonic    string
	APIEndpoint string
	MaxGas      int64
	GasPrice    float64
	UUID        string
}

var _ BluezelleClient = &defaultBluzelleClient{}

type defaultBluzelleClient struct {
	config        *Config
	address       string
	grpcConn      *grpc.ClientConn
	querier       QueryClient
	transactioner TransactionClient
}

func NewBluzelleClient(config *Config) (*defaultBluzelleClient, error) {
	if config.Mnemonic == "" {
		return nil, MISSING_MNEMONIC_ERROR
	}

	grpcConn, err := grpc.Dial(config.APIEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &defaultBluzelleClient{
		config:        config,
		grpcConn:      grpcConn,
		querier:       NewQueryClient(config, grpcConn),
		transactioner: NewTransactionClient(config, grpcConn),
	}, nil
}

func (client *defaultBluzelleClient) Query() QueryClient {
	return client.querier
}

func (client *defaultBluzelleClient) Transaction() TransactionClient {
	return client.transactioner
}

// func (client *defaultBluzelleClient) WithTransactions(ops TransactionOperation...) error {
//
// }

func dialerFunc(ctx context.Context, addr string) (net.Conn, error) {
	return tmnet.Connect(addr)
}
