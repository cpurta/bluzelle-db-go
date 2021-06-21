package bluzelledbgo

import (
	"net/http"

	"github.com/cpurta/bluzelle-db-go/codec/github.com/bluzelle/curium/x/crud/types"
)

type QueryClient interface {
	Read(types.QueryReadRequest) types.QueryReadResponse
	Keys(types.QueryKeysRequest) types.QueryKeysResponse
	MyKeys(types.QueryMyKeysRequest) types.QueryMyKeysResponse
	Count(types.QueryCountRequest) types.QueryCountResponse
	Has(types.QueryHasRequest) types.QueryHasResponse
	Search(types.QuerySearchRequest) types.QuerySearchResponse
	GetNShortestLeases(types.QueryGetNShortestLeasesRequest) types.QueryGetNShortestLeasesResponse
	GetLease(types.QueryGetLeaseRequest) types.QueryGetLeaseResponse
	KeyValues(types.QueryKeyValuesRequest) types.QueryKeyValuesResponse
	File(types.QueryFileRequest) types.QueryFileResponse
}

type defaultQueryClient struct {
	config     *Config
	httpClient *http.Client
}

func NewQueryClient(config *Config, httpClient *http.Client) *defaultQueryClient {
	return &defaultQueryClient{
		config:     config,
		httpClient: httpClient,
	}
}
