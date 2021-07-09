package bluzelledbgo

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/cpurta/bluzelle-db-go/types"
	"github.com/tendermint/tendermint/libs/bytes"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"google.golang.org/protobuf/proto"
)

//go:generate mockgen -source ./transaction.go -destination ./mock_client/mock_transaction.go -package mock_client

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

	NewTransactionCreate(create *pb.MsgCreate) TransactionOperation
	NewTransactionDelete(delete *pb.MsgDelete) TransactionOperation
	NewTransactionDeleteAll(deleteAll *pb.MsgDeleteAll) TransactionOperation
	NewTransactionMultiUpdate(multiUpdate *pb.MsgMultiUpdate) TransactionOperation
	NewTransactionRenewLeasesAll(renewLeasesAll *pb.MsgRenewLeasesAll) TransactionOperation
	NewTransactionRenewLease(renewLease *pb.MsgRenewLease) TransactionOperation
	NewTransactionRename(rename *pb.MsgRename) TransactionOperation
	NewTransactionUpdate(update *pb.MsgUpdate) TransactionOperation
	NewTransactionUpsert(update *pb.MsgUpsert) TransactionOperation
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
		createResponse = &pb.MsgCreateResponse{}
		err            error
	)

	if data, err = proto.Marshal(create); err != nil {
		return nil, err
	}

	b, _ := json.Marshal(bytes.HexBytes(data))

	fmt.Println("Sending msg \"create\" request:", string(b))

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
	var (
		data           []byte
		response       *ctypes.ResultABCIQuery
		value          []byte
		deleteResponse = &pb.MsgDeleteResponse{}
		err            error
	)

	if data, err = proto.Marshal(delete); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, "/bluzelle.curium.crud.Msg/Delete", bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

func (client *defaultTransactionClient) DeleteAll(ctx context.Context, deleteAll *pb.MsgDeleteAll) (*pb.MsgDeleteAllResponse, error) {
	var (
		data              []byte
		response          *ctypes.ResultABCIQuery
		value             []byte
		deleteAllResponse = &pb.MsgDeleteAllResponse{}
		err               error
	)

	if data, err = proto.Marshal(deleteAll); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, "/bluzelle.curium.crud.Msg/DeleteAll", bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, deleteAllResponse); err != nil {
		return nil, err
	}

	return deleteAllResponse, nil
}

func (client *defaultTransactionClient) MultiUpdate(ctx context.Context, multiUpdate *pb.MsgMultiUpdate) (*pb.MsgMultiUpdateResponse, error) {
	var (
		data                []byte
		response            *ctypes.ResultABCIQuery
		value               []byte
		multiUpdateResponse = &pb.MsgMultiUpdateResponse{}
		err                 error
	)

	if data, err = proto.Marshal(multiUpdate); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, "/bluzelle.curium.crud.Msg/MultiUpdate", bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, multiUpdateResponse); err != nil {
		return nil, err
	}

	return multiUpdateResponse, nil
}

func (client *defaultTransactionClient) RenewLeasesAll(ctx context.Context, renewLeasesAll *pb.MsgRenewLeasesAll) (*pb.MsgRenewLeasesAllResponse, error) {
	var (
		data                   []byte
		response               *ctypes.ResultABCIQuery
		value                  []byte
		renewLeasesAllResponse = &pb.MsgRenewLeasesAllResponse{}
		err                    error
	)

	if data, err = proto.Marshal(renewLeasesAll); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, "/bluzelle.curium.crud.Msg/RenewLeasesAll", bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, renewLeasesAllResponse); err != nil {
		return nil, err
	}

	return renewLeasesAllResponse, nil
}

func (client *defaultTransactionClient) RenewLease(ctx context.Context, renewLease *pb.MsgRenewLease) (*pb.MsgRenewLeaseResponse, error) {
	var (
		data               []byte
		response           *ctypes.ResultABCIQuery
		value              []byte
		renewLeaseResponse = &pb.MsgRenewLeaseResponse{}
		err                error
	)

	if data, err = proto.Marshal(renewLease); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, "/bluzelle.curium.crud.Msg/RenewLease", bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, renewLeaseResponse); err != nil {
		return nil, err
	}

	return renewLeaseResponse, nil
}

func (client *defaultTransactionClient) Rename(ctx context.Context, rename *pb.MsgRename) (*pb.MsgRenameResponse, error) {
	var (
		data           []byte
		response       *ctypes.ResultABCIQuery
		value          []byte
		renameResponse = &pb.MsgRenameResponse{}
		err            error
	)

	if data, err = proto.Marshal(rename); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, "/bluzelle.curium.crud.Msg/Rename", bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, renameResponse); err != nil {
		return nil, err
	}

	return renameResponse, nil
}

func (client *defaultTransactionClient) Update(ctx context.Context, update *pb.MsgUpdate) (*pb.MsgUpdateResponse, error) {
	var (
		data           []byte
		response       *ctypes.ResultABCIQuery
		value          []byte
		updateResponse = &pb.MsgUpdateResponse{}
		err            error
	)

	if data, err = proto.Marshal(update); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, "/bluzelle.curium.crud.Msg/Update", bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, updateResponse); err != nil {
		return nil, err
	}

	return updateResponse, nil
}

func (client *defaultTransactionClient) Upsert(ctx context.Context, upsert *pb.MsgUpsert) (*pb.MsgUpsertResponse, error) {
	var (
		data           []byte
		response       *ctypes.ResultABCIQuery
		value          []byte
		upsertResponse = &pb.MsgUpsertResponse{}
		err            error
	)

	if data, err = proto.Marshal(upsert); err != nil {
		return nil, err
	}

	if response, err = client.rpcClient.ABCIQuery(ctx, "/bluzelle.curium.crud.Msg/Upsert", bytes.HexBytes(data)); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, NIL_RESPONSE_RETURNED_ERROR
	}

	value = response.Response.Value

	if err = proto.Unmarshal(value, upsertResponse); err != nil {
		return nil, err
	}

	return upsertResponse, nil
}
