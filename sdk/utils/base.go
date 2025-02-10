package veritechUtils

import (
	"encoding/json"
	"errors"
	"fmt"

	veritechModels "github.com/techpartners-asia/veritech-integration/sdk/models"
)

func GetResponse[T any](sendFunc func(body interface{}) ([]byte, error), body interface{}) (*veritechModels.BaseServiceResponse[T], error) {

	// Call the passed-in function (e.g., `co.Send`)
	res, err := sendFunc(body)
	if err != nil {
		return nil, err
	}

	// return response
	var result veritechModels.BaseServiceResponse[T]
	if err := json.Unmarshal(res, &result); err != nil {
		errors.New("Failed to unmarshal response")
		return nil, err
	}

	if result.Response.Status != "success" {
		fmt.Printf("%v", string(res))
		return nil, errors.New("Response is not success ")
	}

	return &result, nil
}
