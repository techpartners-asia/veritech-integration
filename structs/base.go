package structs

type (
	Base struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		Command      string `json:"command"`
		LanguageCode string `json:"languageCode"`
	}
)
