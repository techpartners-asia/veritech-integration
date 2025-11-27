package warehouseStaffDTO

type (
	IsExistResponseDTO struct {
		IsExist bool `json:"is_exist"`
	}

	StaffResponseDTO struct {
		StoreKeeperKeyID   string `json:"storekeeperkeyid"`
		StoreKeeperKeyCode string `json:"storekeeperkeycode"`
		StoreKeeperName    string `json:"storekeepername"`
	}
)
