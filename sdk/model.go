package veritechSdk

type (
	BaseServiceModel struct {
		Username         string `json:"username"`
		Password         string `json:"password"`
		Command          string `json:"command"`
		LanguageCode     string `json:"languageCode"`
		GroupCode        string `json:"systemmetagroupcode"`
		IgnorePermission string `json:"ignorepermission"`
	}

	// give me base response
	BaseServiceResponse[T any] struct {
		Response struct {
			Status string `json:"status"`
			Result T      `json:"result"`
		}
	}
)
