package policies

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Think-iT-Labs/edc-connector-client/go/edc"
	edchttp "github.com/Think-iT-Labs/edc-connector-client/go/edc/transport/http"
	"github.com/stretchr/testify/assert"
)

func Test_GetPolicy(t *testing.T) {
	authToken := "dummy"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
{
	"id": "1234",
	"createdAt": 1680172087972,
	"policy": {}
}
`)
	}))
	defer svr.Close()

	cfg := edc.NewConfig()

	cfg.Addresses = edc.Addresses{
		Management: &svr.URL,
	}
	httpClient, err := edchttp.NewHTTPClient(&authToken)
	assert.NoError(t, err, "failed to initialize Http Client")
	cfg.HTTPClient = httpClient

	apiClient, err := New(*cfg)
	assert.NoError(t, err, "failed to initialize api client")

	policyId := "1234"
	policy, err := apiClient.GetPolicy(policyId)

	assert.NoError(t, err, "failed to create asset.")
	assert.NotNil(t, policy)
	assert.Equal(t, policy.Id, "1234")
	assert.Equal(t, policy.CreatedAt, int64(1680172087972))
}