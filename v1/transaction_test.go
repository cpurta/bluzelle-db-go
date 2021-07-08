package bluzelledbgo

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	pb "github.com/cpurta/bluzelle-db-go/types"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func generate200Response(log, key string, code int, value []byte) string {
	valueStr := base64.StdEncoding.EncodeToString(value)

	return fmt.Sprintf(`{"error":"","result":{"response":{"log":"%s","height":"0","value":"%s","key":"%s","index":"-1","code":"%d"}},"id":0,"jsonrpc":"2.0"}`, log, valueStr, key, code)
}

func generateProtoResponseData(message proto.Message) []byte {
	data, _ := proto.Marshal(message)

	return data
}

var (
	testCreator = "bluezelleFAKEADDRESS"
	testUUID    = "testUUID"
	testKey     = "fakeKey"
	testValue   = "fakeValue"
)

func TestTransactionCreateHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method     = http.MethodPost
		statusCode = http.StatusOK
		apiURL     = testnetURL
		message    = &pb.MsgCreate{
			Creator: testCreator,
			Uuid:    testUUID,
			Key:     testKey,
			Value:   []byte(testValue),
			Lease: &pb.Lease{
				Hours: uint32(1),
			},
			Metadata: []byte{},
		}
		responseData = generateProtoResponseData(&pb.MsgCreateResponse{})
		config       = &Config{
			APIEndpoint: testnetURL,
		}
		err error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", message.Key, 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	txnClient := NewTransactionClient(config, rpcClient)

	ctx := context.Background()

	_, err = txnClient.Create(ctx, message)

	assert.Nil(t, err)
}

func TestTransactionDeleteHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method     = http.MethodPost
		statusCode = http.StatusOK
		apiURL     = testnetURL
		message    = &pb.MsgDelete{
			Creator: testCreator,
			Uuid:    testUUID,
			Key:     testKey,
		}
		responseData = generateProtoResponseData(&pb.MsgDeleteResponse{})
		config       = &Config{
			APIEndpoint: testnetURL,
		}
		err error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", message.Key, 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	txnClient := NewTransactionClient(config, rpcClient)

	ctx := context.Background()

	_, err = txnClient.Delete(ctx, message)

	assert.Nil(t, err)
}

func TestTransactionDeleteAllHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method     = http.MethodPost
		statusCode = http.StatusOK
		apiURL     = testnetURL
		message    = &pb.MsgDeleteAll{
			Creator: testCreator,
			Uuid:    testUUID,
		}
		responseData = generateProtoResponseData(&pb.MsgDeleteAllResponse{})
		config       = &Config{
			APIEndpoint: testnetURL,
		}
		err error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", "", 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	txnClient := NewTransactionClient(config, rpcClient)

	ctx := context.Background()

	_, err = txnClient.DeleteAll(ctx, message)

	assert.Nil(t, err)
}

func TestTransactionMultiUpdateHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method     = http.MethodPost
		statusCode = http.StatusOK
		apiURL     = testnetURL
		message    = &pb.MsgMultiUpdate{
			Creator: testCreator,
			Uuid:    testUUID,
		}
		responseData = generateProtoResponseData(&pb.MsgMultiUpdateResponse{})
		config       = &Config{
			APIEndpoint: testnetURL,
		}
		err error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", "", 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	txnClient := NewTransactionClient(config, rpcClient)

	ctx := context.Background()

	_, err = txnClient.MultiUpdate(ctx, message)

	assert.Nil(t, err)
}

func TestTransactionRenewLeasesAllHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method     = http.MethodPost
		statusCode = http.StatusOK
		apiURL     = testnetURL
		message    = &pb.MsgRenewLeasesAll{
			Creator: testCreator,
			Uuid:    testUUID,
			Lease: &pb.Lease{
				Hours: uint32(1),
			},
		}
		responseData = generateProtoResponseData(&pb.MsgRenewLeasesAllResponse{})
		config       = &Config{
			APIEndpoint: testnetURL,
		}
		err error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", "", 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	txnClient := NewTransactionClient(config, rpcClient)

	ctx := context.Background()

	_, err = txnClient.RenewLeasesAll(ctx, message)

	assert.Nil(t, err)
}

func TestTransactionRenewLeaseHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method     = http.MethodPost
		statusCode = http.StatusOK
		apiURL     = testnetURL
		message    = &pb.MsgRenewLease{
			Creator: testCreator,
			Uuid:    testUUID,
			Lease: &pb.Lease{
				Hours: uint32(1),
			},
		}
		responseData = generateProtoResponseData(&pb.MsgRenewLeaseResponse{})
		config       = &Config{
			APIEndpoint: testnetURL,
		}
		err error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", "", 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	txnClient := NewTransactionClient(config, rpcClient)

	ctx := context.Background()

	_, err = txnClient.RenewLease(ctx, message)

	assert.Nil(t, err)
}

func TestTransactionRenameHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method     = http.MethodPost
		statusCode = http.StatusOK
		apiURL     = testnetURL
		message    = &pb.MsgRename{
			Creator: testCreator,
			Uuid:    testUUID,
			Key:     testKey,
			NewKey:  "newTestingKey",
		}
		responseData = generateProtoResponseData(&pb.MsgRenameResponse{})
		config       = &Config{
			APIEndpoint: testnetURL,
		}
		err error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", "", 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	txnClient := NewTransactionClient(config, rpcClient)

	ctx := context.Background()

	_, err = txnClient.Rename(ctx, message)

	assert.Nil(t, err)
}

func TestTransactionUpdateHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method     = http.MethodPost
		statusCode = http.StatusOK
		apiURL     = testnetURL
		message    = &pb.MsgUpdate{
			Creator: testCreator,
			Uuid:    testUUID,
			Key:     testKey,
			Value:   []byte(testValue),
			Lease: &pb.Lease{
				Hours: uint32(1),
			},
			Metadata: []byte{},
		}
		responseData = generateProtoResponseData(&pb.MsgUpdateResponse{})
		config       = &Config{
			APIEndpoint: testnetURL,
		}
		err error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", "", 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	txnClient := NewTransactionClient(config, rpcClient)

	ctx := context.Background()

	_, err = txnClient.Update(ctx, message)

	assert.Nil(t, err)
}

func TestTransactionUpsertHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method     = http.MethodPost
		statusCode = http.StatusOK
		apiURL     = testnetURL
		message    = &pb.MsgUpsert{
			Creator: testCreator,
			Uuid:    testUUID,
			Key:     testKey,
			Value:   []byte(testValue),
			Lease: &pb.Lease{
				Hours: uint32(1),
			},
			Metadata: []byte{},
		}
		responseData = generateProtoResponseData(&pb.MsgUpsertResponse{})
		config       = &Config{
			APIEndpoint: testnetURL,
		}
		err error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", "", 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	txnClient := NewTransactionClient(config, rpcClient)

	ctx := context.Background()

	_, err = txnClient.Upsert(ctx, message)

	assert.Nil(t, err)
}
