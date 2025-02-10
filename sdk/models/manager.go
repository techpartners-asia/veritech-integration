package veritechModels

type (
	GetManagerCodeResponse struct {
		StoreKeeperKeyID   string `json:"storekeeperkeyid"`   // Няравын систем дугаар
		StoreKeeperKeyCode string `json:"storekeeperkeycode"` // Нярав код
		StoreKeeperName    string `json:"storekeepername"`    // Нэр
	}
)
