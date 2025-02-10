package veritechSdk

import (
	veritechModels "github.com/techpartners-asia/veritech-integration/sdk/models"
	veritechUtils "github.com/techpartners-asia/veritech-integration/sdk/utils"
)

type (
	TripService interface {
		MakeTrip(input veritechModels.MakeTripBodyModel) (*veritechModels.MakeTripResponse, error)
	}
)

func (co *Sdk) MakeTrip(input veritechModels.MakeTripBodyModel) (*veritechModels.MakeTripResponse, error) {
	co.Command = "PL_MDVIEW_005"
	// co.GroupCode = "imCheckMovementItem_001"

	body := map[string]interface{}{
		"parameters": input,
	}
	body["parameters"].(map[string]interface{})["systemmetagroupcode"] = "imCheckMovementItem_001"
	body["parameters"].(map[string]interface{})["ignorepermission"] = "1"

	res, err := veritechUtils.GetResponse[veritechModels.MakeTripResponse](co.Send, body)
	if err != nil {
		return nil, err
	}

	return &res.Response.Result, nil
}
