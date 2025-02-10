package structs

type (

	// * NOTE * - Inputs
	StoreFillInput struct {
		Base
		Parameters StoreFillParametersInput `json:"parameters"`
	}

	StoreFillParametersInput struct {
		FromStoreKeeperKeyCode string    `json:"fromStoreKeeperKeyCode"`
		ToStoreKeeperKeyCode   string    `json:"toStoreKeeperKeyCode"`
		BookDate               string    `json:"bookDate"`
		Description            string    `json:"description"`
		ImportID               uint      `json:"importId"`
		BookDtl                []BookDtl `json:"bookDtl"`
	}

	BookDtl struct {
		RefItemID   string   `json:"refItemId"`
		TransferQty int      `json:"transferQty"`
		CheckKey    CheckKey `json:"checkKey"`
	}

	CheckKey struct {
		ItemID string `json:"itemId"`
	}

	// * NOTE * - Responses
	StoreFillResponse struct {
		Response FillResponse `json:"response"`
	}
	FillResponse struct {
		Status string     `json:"status"` // error, success
		Result FillResult `json:"result"` // error, success
		Text   string     `json:"text"`   // error, success
	}
	FillResult struct {
		ItemID   string `json:"itemid"`
		ItemCode string `json:"itemcode"`
		ItemName string `json:"itemname"`
	}
)
