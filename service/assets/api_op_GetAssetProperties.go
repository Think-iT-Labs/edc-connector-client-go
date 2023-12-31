package assets

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Think-iT-Labs/edc-connector-client-go/internal"
	"github.com/Think-iT-Labs/edc-connector-client-go/internal/sharedtypes"
)

type GetAssetPropertiesApiResponse struct {
	AssetProperties
	sharedtypes.BaseResponse
}

func (c *Client) GetAssetProperties(assetId string) (*AssetProperties, error) {
	endpoint := fmt.Sprintf("%s/assets/%s", *c.Addresses.Management, assetId)

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

	apiResponse := GetAssetPropertiesApiResponse{}
	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return nil, sdkErrors.FromError(err).FailedTof(internal.ACTION_JSON_UNMARSHAL, response)
	}
	asset := AssetProperties{}
	asset.Id = apiResponse.BaseResponse.Id
	asset.PublicProperties = apiResponse.PublicProperties
	asset.PrivateProperties = apiResponse.PrivateProperties

	return &asset, err
}
