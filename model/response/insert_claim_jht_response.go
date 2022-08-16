package response

type InsertClaimJht struct {
	StatusKirim        string `json:"StatusKirim"`
	KeteranganKirim    string `json:"KeteranganKirim"`
	KodePengajuanKirim string `json:"KodePengajuanKirim"`
}

type InsertClaimJhtResponse struct {
	StatusCode int64            `json:"StatusCode"`
	StatusDesc string           `json:"StatusDesc"`
	Data       []InsertClaimJht `json:"Data"`
}
