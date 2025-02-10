package structs

type (
	ProductInfoInput struct {
		Base
		Paramters ProductInfoParametersInput `json:"parameters"`
	}

	ProductInfoParametersInput struct {
		ItemCode string `json:"itemCode"`
	}

	ProductInfoResponse struct {
		Response ProductInfoResponseTmp `json:"response"`
	}

	ProductInfoResponseTmp struct {
		Status string            `json:"status"` // error, success
		Result ProductInfoResult `json:"result"` // error, success
		Text   string            `json:"text"`   // error, success
	}

	ProductInfoResult struct {
		ItemID   string `json:"itemid"`
		ItemCode string `json:"itemcode"`
		ItemName string `json:"itemname"`
	}
)
