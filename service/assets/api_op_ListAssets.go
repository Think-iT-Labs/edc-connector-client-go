package assets

import (
	"fmt"
	"net/http"

	"github.com/Think-iT-Labs/edc-connector-client-go/common/apivalidator"
	"github.com/Think-iT-Labs/edc-connector-client-go/internal"
	"github.com/Think-iT-Labs/edc-connector-client-go/internal/sharedtypes"
)

type QueryInputPayload struct {
	apivalidator.QueryInput
	sharedtypes.BaseRequest
}

func (c *Client) ListAssets(args ...apivalidator.QueryInput) ([]AssetProperties, error) {
	var queryInputPayload *QueryInputPayload

	if len(args) > 0 {
		queryInputPayload = &QueryInputPayload{}
		queryInputPayload.QueryInput = args[0]
		err := apivalidator.ValidateQueryInput(queryInputPayload.SortOrder)
		if err != nil {
			return nil, err
		}
		queryInputPayload.Context = sharedtypes.EdcContext
	} else {
		queryInputPayload = nil
	}

	endpoint := fmt.Sprintf("%s/assets/request", *c.Addresses.Management)
	assets := []AssetProperties{}
	err := c.HTTPClient.InvokeOperation(internal.InvokeHTTPOperationOptions{
		Method:             http.MethodPost,
		Endpoint:           endpoint,
		RequestPayload:     queryInputPayload,
		ResponsePayload:    &assets,
		ExpectedStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}

	return assets, nil
}
