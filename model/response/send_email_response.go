package response

type SendEmailResponse struct {
	StatusCode string `json:"StatusCode"`
	Message    string `json:"StatusDesc"`
}

type BeforeSendEmailResponse struct {
	StatusCode   string `json:"StatusCode"`
	Message      string `json:"StatusDesc"`
	KanalLayanan string `json:"KanalLayanan"`
}
