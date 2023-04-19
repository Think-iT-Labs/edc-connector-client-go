package contractdefinition

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CreateContractDefinition(cd ContractDefinition) (*CreateContractDefinitionOutput, error) {
	endpoint := fmt.Sprintf("%s/contractdefinitions", *c.Management)
	contractDefinitionApiInput, err := json.Marshal(cd)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while marshaling create contract definition input: %v", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(contractDefinitionApiInput))
	if err != nil {
		return nil, fmt.Errorf("unexpected error while building HTTP request: %v", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while performing POST request to the endpoint %s: %v", endpoint, err)
	}

	defer res.Body.Close()
	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: got %d from %s %s endpoint . Full response : \n %s", res.StatusCode, res.Request.Method, endpoint, response)
	}

	ContractDefinitionOutput := CreateContractDefinitionOutput{}
	err = json.Unmarshal(response, &ContractDefinitionOutput)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling json: %v", err)
	}

	return &ContractDefinitionOutput, err
}
