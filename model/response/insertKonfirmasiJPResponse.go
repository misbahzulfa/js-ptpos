package response

type InsertKonfirmasiJPResponse struct {
	StatusCode int                      `json:"StatusCode"`
	StatusDesc string                   `json:"StatusDesc"`
	Data       []DataInsertKonfirmasiJP `json:"Data"`
}

type DataInsertKonfirmasiJP struct {
	StatusKirim                      string `json:"statusKirim"`
	Keterangan                       string `json:"keterangan"`
	KodePengajuanKonfirmasiJPBerkala string `json:"kodePengajuanKonfirmasiJPBerkala"`
}
