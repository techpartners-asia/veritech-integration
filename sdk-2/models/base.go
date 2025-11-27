package veritechSdkModels

type ResponseStatus string

const (
	ResponseStatusSuccess ResponseStatus = "success"
)

type (
	BaseServiceInput struct {
		Host         string `json:"host"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		LanguageCode string `json:"languageCode"`
	}

	BaseServiceResponse[T any] struct {
		Response struct {
			Status ResponseStatus `json:"status"`
			Result T              `json:"result"`
		} `json:"response"`
	}
)
