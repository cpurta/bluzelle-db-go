package bluzelledbgo

import (
	"net/http"
	"net/url"
)

//go:generate mockgen -source ./client.go -destination ./mock_client/mock.go

// BluezelleClient provides a high level interface to access the Query and Transaction
// rpc methods via Querier and Transactioner interfaces.
type BluezelleClient interface {
	Query() QueryClient
	Transaction() TransactionClient
}

type Config struct {
	mnemonic []string
	apiURL   *url.URL
	maxGas   int64
	gasPrice float64
}

var _ BluezelleClient = &defaultBluezelleClient{}

type defaultBluezelleClient struct {
	config        *Config
	httpClient    *http.Client
	querier       QueryClient
	transactioner TransactionClient
}

func NewBluzelleClient(config *Config, httpClient *http.Client) *defaultBluezelleClient {
	return &defaultBluezelleClient{
		config:     config,
		httpClient: httpClient,
	}
}

func (client *defaultBluezelleClient) Query() QueryClient {
	return client.querier
}

func (client *defaultBluezelleClient) Transaction() TransactionClient {
	return client.transactioner
}
