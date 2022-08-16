package model

type InsertClaimJht struct {
	StatusCode         string `json:"StatusCode"`
	StatusKirim        string `json:"StatusKirim"`
	KeteranganKirim    string `json:"KeteranganKirim"`
	KodePengajuanKirim string `json:"KodePengajuanKirim"`
}

type InsertClaimJhtResponse struct {
	StatusCode string           `json:"StatusCode"`
	StatusDesc string           `json:"StatusDesc"`
	Data       []InsertClaimJht `json:"Data"`
}
