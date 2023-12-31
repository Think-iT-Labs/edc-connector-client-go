package assets

import (
	"fmt"
	"net/http"

	"github.com/Think-iT-Labs/edc-connector-client-go/internal"
	"github.com/Think-iT-Labs/edc-connector-client-go/internal/sharedtypes"
)

type CreateAssetRequestPayload struct {
	sharedtypes.BaseRequest
	AssetProperties `json:"asset"`
	DataAddress     ModifyDataAddressPayload `json:"dataAddress"`
}

func (c *Client) CreateAsset(asset AssetProperties, dataAddress DataAddress) (*sharedtypes.BaseResponse, error) {
	// AssetProperties.Id is offered to the SDK user for convenience.
	// The proper way to set a custom ID for an asset is through the
	// "edc:id" asset property
	if asset.Id != "" {
		asset.PublicProperties["edc:id"] = asset.Id
	}

	requestpayload := CreateAssetRequestPayload{
		BaseRequest: sharedtypes.BaseRequest{
			Context: sharedtypes.EdcContext,
			Id:      asset.Id,
		},
		AssetProperties: asset,
		DataAddress: ModifyDataAddressPayload{
			Type:        getDataAddressType(dataAddress),
			DataAddress: dataAddress,
		},
	}
	endpoint := fmt.Sprintf("%s/assets", *c.Addresses.Management)
	createAssetResponse := sharedtypes.BaseResponse{}
	err := c.HTTPClient.InvokeOperation(internal.InvokeHTTPOperationOptions{
		Method:             http.MethodPost,
		Endpoint:           endpoint,
		RequestPayload:     requestpayload,
		ResponsePayload:    &createAssetResponse,
		ExpectedStatusCode: http.StatusOK,
	})

	if err != nil {
		return nil, err
	}

	return &createAssetResponse, nil
}
