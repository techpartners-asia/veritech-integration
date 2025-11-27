package warehouseStorageDTO

type (
	WarehouseStorageResponseDTO struct {
		WarehouseKeyID   string `json:"warehousekeyid"`
		WarehouseKeyCode string `json:"warehousekeycode"`
		WarehouseName    string `json:"warehousename"`
	}
)
