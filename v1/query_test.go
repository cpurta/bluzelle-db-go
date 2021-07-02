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
	testnetURL       = "https://client.sentry.testnet.private.bluzelle.com:26657"
	QueryResponse200 = `{"jsonrpc":"2.0","id":1234,"result":{"response":{"code":0,"log":"","info":"","index":"0","key":null,"value":"Cgh0ZXN0VVVJRA==","proofOps":null,"height":"95396","codespace":""}}}`
)

func TestQueryCountHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tcs := []struct {
		name             string
		method           string
		statusCode       int
		apiURL           string
		responseBody     string
		expectedResponse *pb.QueryCountResponse
		expectedError    error
	}{
		{
			name:         "200 query response returned",
			method:       http.MethodGet,
			statusCode:   http.StatusInternalServerError,
			apiURL:       testnetURL,
			responseBody: QueryResponse200,
			expectedResponse: &pb.QueryCountResponse{
				Uuid:  "testUUID",
				Count: uint32(0),
			},
			expectedError: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			config := &Config{
				APIEndpoint: testnetURL,
			}

			rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

			assert.Nil(t, err)

			stringResponder := httpmock.NewStringResponder(tc.statusCode, tc.responseBody)

			httpmock.RegisterResponder(tc.method, tc.apiURL, stringResponder)

			queryClient := NewQueryClient(config, rpcClient)

			ctx := context.Background()

			response, err := queryClient.Count(ctx, "testUUID")

			if tc.expectedError != nil && err == nil {
				t.Errorf("expected an error but did not recieve one")
				return
			} else if tc.expectedError != nil && err != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedResponse.Uuid, response.Uuid)
			assert.Equal(t, tc.expectedResponse.Count, response.Count)
		})
	}
}
