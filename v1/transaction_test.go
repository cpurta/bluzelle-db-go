package bluzelledbgo

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	pb "github.com/cpurta/bluzelle-db-go/types"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func generate200TransactionResponse(log, key string, code int, value []byte) string {
	valueStr := string(value)

	return fmt.Sprintf(`{"error":"","result":{"response":{"log":"%s","height":"0","value":"%s","key":"%s","index":"-1","code":"%d"}},"id":0,"jsonrpc":"2.0"}`, log, valueStr, key, code)
}

func TestTransactionCreateHTTPMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tcs := []struct {
		name             string
		method           string
		statusCode       int
		apiURL           string
		createMsg        *pb.MsgCreate
		expectedResponse *pb.MsgCreateResponse
		expectedError    error
	}{
		{
			name:       "200 txn create response returned",
			method:     http.MethodGet,
			statusCode: http.StatusInternalServerError,
			apiURL:     testnetURL,
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
			expectedResponse: &pb.MsgCreateResponse{},
			expectedError:    nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			var (
				config = &Config{
					APIEndpoint: testnetURL,
				}
				responseData []byte
				err          error
			)

			rpcClient, err := client.NewClientFromNode(config.APIEndpoint)

			assert.Nil(t, err)

			if tc.expectedResponse != nil {
				if responseData, err = proto.Marshal(tc.expectedResponse); err != nil {
					t.Errorf("was not expecting an error proto marshalling expected response")
				}
			}

			stringResponder := httpmock.NewStringResponder(tc.statusCode, generate200TransactionResponse("", tc.createMsg.Key, 0, responseData))

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
