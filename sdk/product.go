package veritechSdk

import (
	veritechModels "github.com/techpartners-asia/veritech-integration/sdk/models"
	veritechUtils "github.com/techpartners-asia/veritech-integration/sdk/utils"
)

type ProductService interface {
	GetProductInfo(code string) (*map[string]veritechModels.ProductInfoResponse, error)
}

// * NOTE * : Барааны мэдээлэл авах
func (co *sdk) GetProductInfo(code string) (*map[string]veritechModels.ProductInfoResponse, error) {
	co.Command = "PL_MDVIEW_005"

	body := map[string]interface{}{
		"parameters": map[string]interface{}{
			"systemmetagroupcode": "MTM_ITEM_004",
			"ignorepermission":    "1",
			"itemcode":            code,
		},
	}

	res, err := veritechUtils.GetResponse[map[string]veritechModels.ProductInfoResponse](co.Send, body)
	if err != nil {
		return nil, err
	}

	return &res.Response.Result, nil
}
