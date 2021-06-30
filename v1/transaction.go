package bluzelledbgo

import (
	"context"

	pb "github.com/bluzelle/curium/x/crud/types"
	"github.com/tendermint/tendermint/libs/bytes"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"google.golang.org/protobuf/proto"
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
	config    *Config
	rpcClient *rpchttp.HTTP
}

func NewTransactionClient(config *Config, rpcClient *rpchttp.HTTP) *defaultTransactionClient {
	return &defaultTransactionClient{
		config:    config,
		rpcClient: rpcClient,
	}
}

func (client *defaultTransactionClient) Create(ctx context.Context, create *pb.MsgCreate) (*pb.MsgCreateResponse, error) {
	var (
		data           []byte
		response       *ctypes.ResultABCIQuery
		value          []byte
		createResponse *pb.MsgCreateResponse
		err            error
	)

	if data, err = proto.Marshal(create); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, "/bluzelle.curium.crud.Msg/Create", bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, createResponse); err != nil {
		return nil, err
	}

	return createResponse, nil
}

func (client *defaultTransactionClient) Delete(ctx context.Context, delete *pb.MsgDelete) (*pb.MsgDeleteResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultTransactionClient) DeleteAll(ctx context.Context, deleteAll *pb.MsgDeleteAll) (*pb.MsgDeleteAllResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultTransactionClient) MultiUpdate(ctx context.Context, multiUpdate *pb.MsgMultiUpdate) (*pb.MsgMultiUpdateResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultTransactionClient) RenewLeasesAll(ctx context.Context, renewLeasesAll *pb.MsgRenewLeasesAll) (*pb.MsgRenewLeasesAllResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultTransactionClient) RenewLease(ctx context.Context, renewLease *pb.MsgRenewLease) (*pb.MsgRenewLeaseResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultTransactionClient) Rename(ctx context.Context, rename *pb.MsgRename) (*pb.MsgRenameResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultTransactionClient) Update(ctx context.Context, update *pb.MsgUpdate) (*pb.MsgUpdateResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}

func (client *defaultTransactionClient) Upsert(ctx context.Context, upsert *pb.MsgUpsert) (*pb.MsgUpsertResponse, error) {
	// TODO: implement
	return nil, NOT_IMPLEMENTED_ERROR
}
