package bluzelledbgo

import (
	"context"
	"net/http"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	pb "github.com/cpurta/bluzelle-db-go/types"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	testnetURL = "https://client.sentry.testnet.private.bluzelle.com:26657"
)

func TestQueryCountHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method       = http.MethodPost
		statusCode   = http.StatusOK
		apiURL       = testnetURL
		responseData = `{"jsonrpc":"2.0","id":1234,"result":{"response":{"code":0,"log":"","info":"","index":"0","key":null,"value":"Cgh0ZXN0VVVJRA==","proofOps":null,"height":"204818","codespace":""}}}`
		config       = &Config{
			APIEndpoint: testnetURL,
			UUID:        testUUID,
		}
		resp *pb.QueryCountResponse
		err  error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, responseData)

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	queryClient := NewQueryClient(config, rpcClient)

	ctx := context.Background()

	resp, err = queryClient.Count(ctx)

	assert.Nil(t, err)
	assert.Equal(t, testUUID, resp.Uuid)
	assert.Equal(t, uint32(0), resp.Count)
}

func TestGetLeaseHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method       = http.MethodPost
		statusCode   = http.StatusOK
		apiURL       = testnetURL
		responseData = generateProtoResponseData(&pb.QueryGetLeaseResponse{
			Uuid:    testUUID,
			Key:     testKey,
			Seconds: uint32(60),
		})
		config = &Config{
			APIEndpoint: testnetURL,
			UUID:        testUUID,
		}
		resp *pb.QueryGetLeaseResponse
		err  error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", testKey, 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	queryClient := NewQueryClient(config, rpcClient)

	ctx := context.Background()

	resp, err = queryClient.GetLease(ctx, testKey)

	assert.Nil(t, err)
	assert.Equal(t, testUUID, resp.Uuid)
	assert.Equal(t, testKey, resp.Key)
	assert.Equal(t, uint32(60), resp.Seconds)
}

func TestGetNShortestLeasesHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var (
		method       = http.MethodPost
		statusCode   = http.StatusOK
		apiURL       = testnetURL
		responseData = generateProtoResponseData(&pb.QueryGetNShortestLeasesResponse{
			Uuid: testUUID,
			KeyLeases: []*pb.KeyLease{
				&pb.KeyLease{
					Key:     testKey,
					Seconds: uint32(60),
				},
			},
		})
		config = &Config{
			APIEndpoint: testnetURL,
			UUID:        testUUID,
		}
		resp *pb.QueryGetNShortestLeasesResponse
		err  error
	)

	rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

	assert.Nil(t, err)

	stringResponder := httpmock.NewStringResponder(statusCode, generate200Response("", testKey, 0, responseData))

	httpmock.RegisterResponder(method, apiURL, stringResponder)

	queryClient := NewQueryClient(config, rpcClient)

	ctx := context.Background()

	resp, err = queryClient.GetNShortestLeases(ctx, 10)

	assert.Nil(t, err)
	assert.Equal(t, testUUID, resp.Uuid)
	assert.Equal(t, 1, len(resp.KeyLeases))
}
