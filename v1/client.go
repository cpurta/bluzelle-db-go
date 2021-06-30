package bluzelledbgo

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
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
	rpcClient     *rpchttp.HTTP
	querier       QueryClient
	transactioner TransactionClient
}

func NewBluzelleClient(config *Config) (*defaultBluzelleClient, error) {
	if config.Mnemonic == "" {
		return nil, MISSING_MNEMONIC_ERROR
	}

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to connect rpc websocket: %s", err.Error())
	}

	return &defaultBluzelleClient{
		config:        config,
		rpcClient:     rpcClient,
		querier:       NewQueryClient(config, rpcClient),
		transactioner: NewTransactionClient(config, rpcClient),
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
