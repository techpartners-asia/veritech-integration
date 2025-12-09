package reportFinanceDTO

type (
	ListResponseDTO struct {
		StoreKeeperKeyCode string                         `json:"storekeeperkeycode"`
		BookDate           string                         `json:"bookdate"`
		DtlSize            int                            `json:"dtlsize"`
		Dtl                map[string]ListResponseItemDTO `json:"dtl"`
	}

	ListResponseItemDTO struct {
		StoreKeeperKeyCode string  `json:"storekeeperkeycode"`
		BookDate           string  `json:"bookdate"`
		ItemID             string  `json:"itemid"`
		ItemCode           string  `json:"itemcode"`
		InQty              float64 `json:"inqty"`
		OutQty             float64 `json:"outqty"`
		TQty               float64 `json:"tqty"`
	}

	GetInitialQtyResponseDTO struct {
		StoreKeeperKeyCode string                                  `json:"storekeeperkeycode"`
		DtlSize            int                                     `json:"dtlsize"`
		Dtl                map[string]GetInitialQtyResponseItemDTO `json:"dtl"`
	}
	GetInitialQtyResponseItemDTO struct {
		StoreKeeperKeyCode string `json:"storekeeperkeycode"`
		ItemID             string `json:"itemid"`
		ItemCode           string `json:"itemcode"`
		EndQty             int    `json:"end_qty"`
	}
)
