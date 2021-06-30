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
	TransactionCreateResponse200 = `{"error":"","result":{"response":{"log":"","height":"0","value":"","key":"","index":"-1","code":"0"}},"id":0,"jsonrpc":"2.0"}`
)

func TestTransactionCreateHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tcs := []struct {
		name             string
		method           string
		statusCode       int
		apiURL           string
		createMsg        *pb.MsgCreate
		responseBody     string
		expectedResponse *pb.QueryCountResponse
		expectedError    error
	}{
		{
			name:       "200 txn create response returned",
			method:     http.MethodGet,
			statusCode: http.StatusInternalServerError,
			apiURL:     testnetURL + "/abci_query",
			createMsg: &pb.MsgCreate{
				Creator: "bluezelleFAKEADDRESS",
				Uuid:    "testUUID",
				Key:     "fakeKey",
				Value:   []byte("fakeValue"),
				Lease: &pb.Lease{
					Hours: uint32(1),
				},
				Metadata: []byte{},
			},
			responseBody:  TransactionCreateResponse200,
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

			txnClient := NewTransactionClient(config, rpcClient)

			ctx := context.Background()

			_, err = txnClient.Create(ctx, tc.createMsg)

			if tc.expectedError != nil && err == nil {
				t.Errorf("expected an error but did not recieve one")
				return
			} else if tc.expectedError != nil && err != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
		})
	}
}
