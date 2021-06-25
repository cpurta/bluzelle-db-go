package bluzelledbgo

import (
	"context"

	pb "github.com/cpurta/bluzelle-db-go/types"
	"google.golang.org/grpc"
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
	config   *Config
	pbClient pb.QueryClient
}

func NewQueryClient(config *Config, grpcConn *grpc.ClientConn) *defaultQueryClient {
	pbClient := pb.NewQueryClient(grpcConn)

	return &defaultQueryClient{
		config:   config,
		pbClient: pbClient,
	}
}

func (client *defaultQueryClient) Count(ctx context.Context, uuid string) (*pb.QueryCountResponse, error) {
	return client.pbClient.Count(ctx, &pb.QueryCountRequest{
		Uuid: uuid,
	})
}

func (client *defaultQueryClient) GetLease(ctx context.Context, uuid, key string) (*pb.QueryGetLeaseResponse, error) {
	return client.pbClient.GetLease(ctx, &pb.QueryGetLeaseRequest{
		Uuid: uuid,
		Key:  key,
	})
}

func (client *defaultQueryClient) GetNShortestLeases(ctx context.Context, uuid string, number int) (*pb.QueryGetNShortestLeasesResponse, error) {
	return client.pbClient.GetNShortestLeases(ctx, &pb.QueryGetNShortestLeasesRequest{
		Uuid: uuid,
		Num:  uint32(number),
	})
}

func (client *defaultQueryClient) Has(ctx context.Context, uuid, key string) (*pb.QueryHasResponse, error) {
	return client.pbClient.Has(ctx, &pb.QueryHasRequest{
		Uuid: uuid,
		Key:  key,
	})
}

func (client *defaultQueryClient) Keys(ctx context.Context, uuid, startKey string, limit int64) (*pb.QueryKeysResponse, error) {
	return client.pbClient.Keys(ctx, &pb.QueryKeysRequest{
		Uuid: uuid,
		Pagination: &pb.PagingRequest{
			StartKey: startKey,
			Limit:    uint64(limit),
		},
	})
}

func (client *defaultQueryClient) KeyValues(ctx context.Context, uuid, startKey string, limit int64) (*pb.QueryKeyValuesResponse, error) {
	return client.pbClient.KeyValues(ctx, &pb.QueryKeyValuesRequest{
		Uuid: uuid,
		Pagination: &pb.PagingRequest{
			StartKey: startKey,
			Limit:    uint64(limit),
		},
	})
}

func (client *defaultQueryClient) MyKeys(ctx context.Context, uuid, address string) (*pb.QueryMyKeysResponse, error) {
	return client.pbClient.MyKeys(ctx, &pb.QueryMyKeysRequest{
		Uuid:    uuid,
		Address: address,
	})
}

func (client *defaultQueryClient) Read(ctx context.Context, uuid, key string) (*pb.QueryReadResponse, error) {
	return client.pbClient.Read(ctx, &pb.QueryReadRequest{
		Uuid: uuid,
		Key:  key,
	})
}

func (client *defaultQueryClient) Search(ctx context.Context, uuid, searchString, startKey string, limit int64) (*pb.QuerySearchResponse, error) {
	return client.pbClient.Search(ctx, &pb.QuerySearchRequest{
		Uuid:         uuid,
		SearchString: searchString,
		Pagination: &pb.PagingRequest{
			StartKey: startKey,
			Limit:    uint64(limit),
		},
	})
}
