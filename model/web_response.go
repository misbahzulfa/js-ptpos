package model

type WebResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"statusDesc"`
}
