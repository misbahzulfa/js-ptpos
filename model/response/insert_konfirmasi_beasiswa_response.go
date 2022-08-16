package response

type InsertKonfirmasiBeasiswaResponse struct {
	StatusCode               string                     `json:"statusCode"`
	StatusDesc               string                     `json:"statusDesc"`
	InsertKonfirmasiBeasiswa []InsertKonfirmasiBeasiswa `json:"dataInsertKonfirmasiBeasiswa"`
}

type InsertKonfirmasiBeasiswa struct {
	KodePengajuanKonfirmasi string `json:"kodePengajuanKonfirmasi"`
	StatusKirim             string `json:"statusKirim"`
	KeteranganSukses        string `json:"keteranganSukses"`
}
