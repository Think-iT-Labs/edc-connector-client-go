package assets

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Think-iT-Labs/edc-connector-client-go/internal"
)

func (c *Client) GetAssetDataAddress(assetId string) (DataAddress, error) {
	endpoint := fmt.Sprintf("%s/assets/%s/dataaddress", *c.Addresses.Management, assetId)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, sdkErrors.FromError(err).FailedTo(internal.ACTION_HTTP_BUILD_REQUEST)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, sdkErrors.FromError(err).FailedTo(internal.ACTION_HTTP_DO_REQUEST)
	}

	defer res.Body.Close()
	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, sdkErrors.FromError(err).FailedTo(internal.ACTION_HTTP_READ_BYTES)
	}

	if res.StatusCode != http.StatusOK {
		return nil, sdkErrors.FromError(internal.ParseConnectorApiError(response)).Error(internal.ERROR_API_ERROR)
	}

	return unmarshalToConcreteDataAddress(response)
}
