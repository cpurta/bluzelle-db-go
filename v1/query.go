package bluzelledbgo

import (
	"context"

	pb "github.com/cpurta/bluzelle-db-go/types"
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
	MyKeys(ctx context.Context, uuid, address, startKey string, limit int64) (*pb.QueryMyKeysResponse, error)
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
		path               = "/bluzelle.curium.crud.Query/Count"
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
	var (
		queryGetLeaseRequest = &pb.QueryGetLeaseRequest{
			Uuid: uuid,
		}
		path                  = "/bluzelle.curium.crud.Query/GetLease"
		data                  []byte
		response              *ctypes.ResultABCIQuery
		value                 []byte
		queryGetLeaseResponse = &pb.QueryGetLeaseResponse{}
		err                   error
	)

	if data, err = proto.Marshal(queryGetLeaseRequest); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, path, bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, queryGetLeaseResponse); err != nil {
		return nil, err
	}

	return queryGetLeaseResponse, nil
}

func (client *defaultQueryClient) GetNShortestLeases(ctx context.Context, uuid string, number int) (*pb.QueryGetNShortestLeasesResponse, error) {
	var (
		queryGetNShortestLeasesRequest = &pb.QueryGetNShortestLeasesRequest{
			Uuid: uuid,
			Num:  uint32(number),
		}
		path                            = "/bluzelle.curium.crud.Query/GetNShortestLeases"
		data                            []byte
		response                        *ctypes.ResultABCIQuery
		value                           []byte
		queryGetNShortestLeasesResponse = &pb.QueryGetNShortestLeasesResponse{}
		err                             error
	)

	if data, err = proto.Marshal(queryGetNShortestLeasesRequest); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, path, bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, queryGetNShortestLeasesResponse); err != nil {
		return nil, err
	}

	return queryGetNShortestLeasesResponse, nil
}

func (client *defaultQueryClient) Has(ctx context.Context, uuid, key string) (*pb.QueryHasResponse, error) {
	var (
		queryHasRequest = &pb.QueryHasRequest{
			Uuid: uuid,
			Key:  key,
		}
		path             = "/bluzelle.curium.crud.Query/Has"
		data             []byte
		response         *ctypes.ResultABCIQuery
		value            []byte
		queryHasResponse = &pb.QueryHasResponse{}
		err              error
	)

	if data, err = proto.Marshal(queryHasRequest); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, path, bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, queryHasResponse); err != nil {
		return nil, err
	}

	return queryHasResponse, nil
}

func (client *defaultQueryClient) Keys(ctx context.Context, uuid, startKey string, limit int64) (*pb.QueryKeysResponse, error) {
	var (
		queryKeysRequest = &pb.QueryKeysRequest{
			Uuid: uuid,
			Pagination: &pb.PagingRequest{
				StartKey: startKey,
				Limit:    uint64(limit),
			},
		}
		path              = "/bluzelle.curium.crud.Query/Keys"
		data              []byte
		response          *ctypes.ResultABCIQuery
		value             []byte
		queryKeysResponse = &pb.QueryKeysResponse{}
		err               error
	)

	if data, err = proto.Marshal(queryKeysRequest); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, path, bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, queryKeysResponse); err != nil {
		return nil, err
	}

	return queryKeysResponse, nil
}

func (client *defaultQueryClient) KeyValues(ctx context.Context, uuid, startKey string, limit int64) (*pb.QueryKeyValuesResponse, error) {
	var (
		queryKeyValuesRequest = &pb.QueryKeyValuesRequest{
			Uuid: uuid,
			Pagination: &pb.PagingRequest{
				StartKey: startKey,
				Limit:    uint64(limit),
			},
		}
		path                   = "/bluzelle.curium.crud.Query/KeyValues"
		data                   []byte
		response               *ctypes.ResultABCIQuery
		value                  []byte
		queryKeyValuesResponse = &pb.QueryKeyValuesResponse{}
		err                    error
	)

	if data, err = proto.Marshal(queryKeyValuesRequest); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, path, bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, queryKeyValuesResponse); err != nil {
		return nil, err
	}

	return queryKeyValuesResponse, nil
}

func (client *defaultQueryClient) MyKeys(ctx context.Context, uuid, address, startKey string, limit int64) (*pb.QueryMyKeysResponse, error) {
	var (
		queryMyKeysRequest = &pb.QueryMyKeysRequest{
			Uuid:    uuid,
			Address: address,
			Pagination: &pb.PagingRequest{
				StartKey: startKey,
				Limit:    uint64(limit),
			},
		}
		path                = "/bluzelle.curium.crud.Query/MyKeys"
		data                []byte
		response            *ctypes.ResultABCIQuery
		value               []byte
		queryMyKeysResponse = &pb.QueryMyKeysResponse{}
		err                 error
	)

	if data, err = proto.Marshal(queryMyKeysRequest); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, path, bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, queryMyKeysResponse); err != nil {
		return nil, err
	}

	return queryMyKeysResponse, nil
}

func (client *defaultQueryClient) Read(ctx context.Context, uuid, key string) (*pb.QueryReadResponse, error) {
	var (
		queryReadRequest = &pb.QueryReadRequest{
			Uuid: uuid,
			Key:  key,
		}
		path              = "/bluzelle.curium.crud.Query/Read"
		data              []byte
		response          *ctypes.ResultABCIQuery
		value             []byte
		queryReadResponse = &pb.QueryReadResponse{}
		err               error
	)

	if data, err = proto.Marshal(queryReadRequest); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, path, bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, queryReadResponse); err != nil {
		return nil, err
	}

	return queryReadResponse, nil
}

func (client *defaultQueryClient) Search(ctx context.Context, uuid, searchString, startKey string, limit int64) (*pb.QuerySearchResponse, error) {
	var (
		querySearchRequest = &pb.QuerySearchRequest{
			Uuid:         uuid,
			SearchString: searchString,
			Pagination: &pb.PagingRequest{
				StartKey: startKey,
				Limit:    uint64(limit),
			},
		}
		path                = "/bluzelle.curium.crud.Query/Search"
		data                []byte
		response            *ctypes.ResultABCIQuery
		value               []byte
		querySearchResponse = &pb.QuerySearchResponse{}
		err                 error
	)

	if data, err = proto.Marshal(querySearchRequest); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, path, bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, querySearchResponse); err != nil {
		return nil, err
	}

	return querySearchResponse, nil
}
