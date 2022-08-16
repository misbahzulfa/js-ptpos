package model

type InsertPengajuanJHT struct {
	StatusCode         int    `json:"StatusCode"`
	StatusKirim        string `json:"StatusKirim"`
	KeteranganKirim    string `json:"KeteranganKirim"`
	KodePengajuanKlaim string `json:"KodePengajuanKlaim"`
}
type InsertPengajuanJHTResponse struct {
	StatusCode int                       `json:"StatusCode"`
	Data       []DaftarDokumenSebabKlaim `json:"Data"`
}
