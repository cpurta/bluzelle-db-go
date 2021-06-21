package bluzelledbgo

import (
	"net/http"

	"github.com/cpurta/bluzelle-db-go/codec/github.com/bluzelle/curium/x/crud/types"
)

type TransactionClient interface {
	Count(types.MsgCount) types.MsgCountResponse
	RenewLeasesAll(types.MsgRenewLeasesAll) types.MsgRenewLeasesAllResponse
	RenewLease(types.MsgRenewLease) types.MsgRenewLeaseResponse
	GetNShortestLeases(types.MsgGetNShortestLeases) types.MsgGetNShortestLeasesResponse
	Keys(types.MsgKeys) types.MsgKeysResponse
	Rename(types.MsgRename) types.MsgRenameResponse
	MultiUpdate(types.MsgMultiUpdate) types.MsgMultiUpdateResponse
	DeleteAll(types.MsgDeleteAll) types.MsgDeleteAllResponse
	KeyValues(types.MsgKeyValues) types.MsgKeyValuesResponse
	Has(types.MsgHas) types.MsgHasResponse
	GetLease(types.MsgGetLease) types.MsgGetLeaseResponse
	Read(types.MsgRead) types.MsgReadResponse
	Upsert(types.MsgUpsert) types.MsgUpsertResponse
	Create(types.MsgCreate) types.MsgCreateResponse
	Update(types.MsgUpdate) types.MsgUpdateResponse
	Delete(types.MsgDelete) types.MsgDeleteResponse
}

type defaultTransactionClient struct {
	config     *Config
	httpClient *http.Client
}

func NewTransactionClient(config *Config, httpClient *http.Client) *defaultTransactionClient {
	return &defaultTransactionClient{
		config:     config,
		httpClient: httpClient,
	}
}
