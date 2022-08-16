package entity

type InsertPengajuanJHTEntity struct {
	StatusCode         int    `json:"StatusCode"`
	StatusKirim        string `json:"StatusKirim"`
	KeteranganKirim    string `json:"KeteranganKirim"`
	KodePengajuanKlaim string `json:"KodePengajuanKlaim"`
}
