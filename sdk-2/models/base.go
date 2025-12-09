package veritechSdkModels

type ResponseStatus string

const (
	ResponseStatusSuccess ResponseStatus = "success"
	ReponseStatusError    ResponseStatus = "error"
)

type (
	BaseServiceInput struct {
		Host         string `json:"host"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		StoreName    string `json:"storeName"`
		StoreCode    string `json:"storeCode"`
		LanguageCode string `json:"languageCode"`
	}

	BaseServiceResponse[T any] struct {
		Response struct {
			Status ResponseStatus `json:"status"`
			Result T              `json:"result"`
			Text   string         `json:"text"`
			Code   string         `json:"code"`
		} `json:"response"`
	}
)
