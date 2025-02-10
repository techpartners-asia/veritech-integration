package veritechSdk

import (
	veritechModels "github.com/techpartners-asia/veritech-integration/sdk/models"
	veritechUtils "github.com/techpartners-asia/veritech-integration/sdk/utils"
)

type StaffManagerService interface {
	GetManagerCode(code string) string
	GetStaffCode(code string) string
}

func (co *Sdk) GetManagerCode(code string) string {

	co.Command = "mtm_sales_skk_code_004"

	body := map[string]interface{}{
		"parameters": map[string]interface{}{
			"filterStoreKeeperKeyCode": code,
		},
	}

	res, err := veritechUtils.GetResponse[veritechModels.GetManagerCodeResponse](co.Send, body)
	if err != nil {
		return ""
	}

	return res.Response.Result.StoreKeeperKeyCode
}

func (co *Sdk) GetStaffCode(code string) string {
	co.Command = "mtm_sales_skk_code_004"

	body := map[string]interface{}{
		"parameters": map[string]interface{}{
			"filterStoreKeeperKeyCode": code,
		},
	}

	res, err := veritechUtils.GetResponse[veritechModels.GetManagerCodeResponse](co.Send, body)
	if err != nil {
		return ""
	}

	return res.Response.Result.StoreKeeperKeyCode
}
