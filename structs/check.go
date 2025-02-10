package structs

type (
	StaffMachineCheckInput struct {
		Code    string `json:"code"`
		Command string `json:"command"`
	}

	CheckRequest struct {
		Request *RequestBody `json:"request"`
	}

	RequestBody struct {
		Command    string      `json:"command"`
		Parameters interface{} `json:"parameters"`
		Username   string      `json:"username"`
		Password   string      `json:"password"`
	}

	StaffMachineCheckResponse struct {
		Response struct {
			Status string      `json:"status"`
			Result interface{} `json:"result"`
		} `json:"response"`
	}

	TakeHTBeginQtyResponse struct {
		Response struct {
			Status string                `json:"status"`
			Result *TakeHTBeginQtyResult `json:"result"`
		} `json:"response"`
	}

	TakeHTBeginQtyResult struct {
		StorekeeperKeyCode string                `json:"storekeeperkeycode"`
		DtlSize            string                `json:"dtl_size"`
		Dtl                map[string]ProductDtl `json:"dtl"`
	}

	ProductDtl struct {
		StorekeeperKeyCode string `json:"storekeeperkeycode"`
		ItemID             string `json:"itemid"`
		ItemCode           string `json:"itemcode"`
		EndQty             string `json:"end_qty"`
	}

	StaffWarehouseIncomeOutcome struct {
		Response struct {
			Status string                             `json:"status"`
			Result *StaffWarehouseIncomeOutcomeResult `json:"result"`
		} `json:"response"`
	}

	StaffWarehouseIncomeOutcomeResult struct {
		StorekeeperKeyCode string                              `json:"storekeeperkeycode"`
		DtlSize            string                              `json:"dtl_size"`
		Dtl                map[string]StaffWarehouseChangesDtl `json:"dtl"`
	}

	StaffWarehouseChangesDtl struct {
		StorekeeperKeyCode string `json:"storekeeperkeycode"`
		ItemID             string `json:"itemid"`
		ItemCode           string `json:"itemcode"`
		Inqty              string `json:"inqty"`
		Outqty             string `json:"outqty"`
		Tqty               string `json:"tqty"`
	}
)
