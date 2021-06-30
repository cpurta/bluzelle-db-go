package bluzelledbgo

import (
	"context"
	"fmt"

	pb "github.com/bluzelle/curium/x/crud/types"
	"github.com/tendermint/tendermint/libs/bytes"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"google.golang.org/protobuf/proto"
)

type QueryClient interface {
	Count(ctx context.Context, uuid string) (*pb.QueryCountResponse, error)
	GetLease(ctx context.Context, uuid, key string) (*pb.QueryGetLeaseResponse, error)
	GetNShortestLeases(ctx context.Context, uuid string, number int) (*pb.QueryGetNShortestLeasesResponse, error)
	Has(ctx context.Context, uuid, key string) (*pb.QueryHasResponse, error)
	Keys(ctx context.Context, uuid, startKey string, limit int64) (*pb.QueryKeysResponse, error)
	KeyValues(ctx context.Context, uuid, startKey string, limit int64) (*pb.QueryKeyValuesResponse, error)
	MyKeys(ctx context.Context, uuid, address string) (*pb.QueryMyKeysResponse, error)
	Read(ctx context.Context, uuid, key string) (*pb.QueryReadResponse, error)
	Search(ctx context.Context, uuid, searchString, startKey string, limit int64) (*pb.QuerySearchResponse, error)
}

var _ QueryClient = &defaultQueryClient{}

type defaultQueryClient struct {
	config    *Config
	rpcClient *rpchttp.HTTP
}

func NewQueryClient(config *Config, rpcClient *rpchttp.HTTP) *defaultQueryClient {
	return &defaultQueryClient{
		config:    config,
		rpcClient: rpcClient,
	}
}

func (client *defaultQueryClient) Count(ctx context.Context, uuid string) (*pb.QueryCountResponse, error) {
	var (
		queryCountRequest = &pb.QueryCountRequest{
			Uuid: uuid,
		}
		path               = fmt.Sprintf("custom/%s/list-CrudValue", pb.QuerierRoute)
		data               []byte
		response           *ctypes.ResultABCIQuery
		value              []byte
		queryCountResponse = &pb.QueryCountResponse{}
		err                error
	)

	if data, err = proto.Marshal(queryCountRequest); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, path, bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, queryCountResponse); err != nil {
		return nil, err
	}

	return queryCountResponse, nil
}

func (client *defaultQueryClient) GetLease(ctx context.Context, uuid, key string) (*pb.QueryGetLeaseResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultQueryClient) GetNShortestLeases(ctx context.Context, uuid string, number int) (*pb.QueryGetNShortestLeasesResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultQueryClient) Has(ctx context.Context, uuid, key string) (*pb.QueryHasResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultQueryClient) Keys(ctx context.Context, uuid, startKey string, limit int64) (*pb.QueryKeysResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultQueryClient) KeyValues(ctx context.Context, uuid, startKey string, limit int64) (*pb.QueryKeyValuesResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultQueryClient) MyKeys(ctx context.Context, uuid, address string) (*pb.QueryMyKeysResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultQueryClient) Read(ctx context.Context, uuid, key string) (*pb.QueryReadResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultQueryClient) Search(ctx context.Context, uuid, searchString, startKey string, limit int64) (*pb.QuerySearchResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}
