package entity

type InsertKonfirmasiBeasiswaEntity struct {
	KodePengajuanKonfirmasi string `json:"kodePengajuanKonfirmasi"`
	StatusKirim             string `json:"statusKirim"`
	KeteranganSukses        string `json:"keteranganSukses"`
}
