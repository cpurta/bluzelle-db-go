package bluzelledbgo

import (
	"context"

	pb "github.com/cpurta/bluzelle-db-go/types"
	"google.golang.org/grpc"
)

type TransactionClient interface {
	Create(ctx context.Context, create *pb.MsgCreate) (*pb.MsgCreateResponse, error)
	Delete(ctx context.Context, delete *pb.MsgDelete) (*pb.MsgDeleteResponse, error)
	DeleteAll(ctx context.Context, deleteAll *pb.MsgDeleteAll) (*pb.MsgDeleteAllResponse, error)
	MultiUpdate(ctx context.Context, multiUpdate *pb.MsgMultiUpdate) (*pb.MsgMultiUpdateResponse, error)
	RenewLeasesAll(ctx context.Context, renewLeasesAll *pb.MsgRenewLeasesAll) (*pb.MsgRenewLeasesAllResponse, error)
	RenewLease(ctx context.Context, renewLease *pb.MsgRenewLease) (*pb.MsgRenewLeaseResponse, error)
	Rename(ctx context.Context, rename *pb.MsgRename) (*pb.MsgRenameResponse, error)
	Update(ctx context.Context, update *pb.MsgUpdate) (*pb.MsgUpdateResponse, error)
	Upsert(ctx context.Context, upsert *pb.MsgUpsert) (*pb.MsgUpsertResponse, error)
}

var _ TransactionClient = &defaultTransactionClient{}

type defaultTransactionClient struct {
	config   *Config
	pbClient pb.MsgClient
}

func NewTransactionClient(config *Config, grpcConn *grpc.ClientConn) *defaultTransactionClient {
	pbClient := pb.NewMsgClient(grpcConn)

	return &defaultTransactionClient{
		config:   config,
		pbClient: pbClient,
	}
}

func (client *defaultTransactionClient) Create(ctx context.Context, create *pb.MsgCreate) (*pb.MsgCreateResponse, error) {
	return client.pbClient.Create(ctx, create)
}

func (client *defaultTransactionClient) Delete(ctx context.Context, delete *pb.MsgDelete) (*pb.MsgDeleteResponse, error) {
	return client.pbClient.Delete(ctx, delete)
}

func (client *defaultTransactionClient) DeleteAll(ctx context.Context, deleteAll *pb.MsgDeleteAll) (*pb.MsgDeleteAllResponse, error) {
	return client.pbClient.DeleteAll(ctx, deleteAll)
}

func (client *defaultTransactionClient) MultiUpdate(ctx context.Context, multiUpdate *pb.MsgMultiUpdate) (*pb.MsgMultiUpdateResponse, error) {
	return client.pbClient.MultiUpdate(ctx, multiUpdate)
}

func (client *defaultTransactionClient) RenewLeasesAll(ctx context.Context, renewLeasesAll *pb.MsgRenewLeasesAll) (*pb.MsgRenewLeasesAllResponse, error) {
	return client.pbClient.RenewLeasesAll(ctx, renewLeasesAll)
}

func (client *defaultTransactionClient) RenewLease(ctx context.Context, renewLease *pb.MsgRenewLease) (*pb.MsgRenewLeaseResponse, error) {
	return client.pbClient.RenewLease(ctx, renewLease)
}

func (client *defaultTransactionClient) Rename(ctx context.Context, rename *pb.MsgRename) (*pb.MsgRenameResponse, error) {
	return client.pbClient.Rename(ctx, rename)
}

func (client *defaultTransactionClient) Update(ctx context.Context, update *pb.MsgUpdate) (*pb.MsgUpdateResponse, error) {
	return client.pbClient.Update(ctx, update)
}

func (client *defaultTransactionClient) Upsert(ctx context.Context, upsert *pb.MsgUpsert) (*pb.MsgUpsertResponse, error) {
	return client.pbClient.Upsert(ctx, upsert)
}
