package bluzelledbgo

import (
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
	Mnemonic   []string
	ApiAddress string
	MaxGas     int64
	GasPrice   float64
}

var _ BluezelleClient = &defaultBluezelleClient{}

type defaultBluezelleClient struct {
	config        *Config
	grpcConn      *grpc.ClientConn
	querier       QueryClient
	transactioner TransactionClient
}

func NewBluzelleClient(config *Config) (*defaultBluezelleClient, error) {
	grpcConn, err := grpc.Dial(config.ApiAddress, grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &defaultBluezelleClient{
		config:   config,
		grpcConn: grpcConn,
	}, nil
}

func (client *defaultBluezelleClient) Query() QueryClient {
	return client.querier
}

func (client *defaultBluezelleClient) Transaction() TransactionClient {
	return client.transactioner
}
