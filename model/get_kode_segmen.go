package model

type DaftarSegmen struct {
	Kode       string `json:"Kode"`
	Keterangan string `json:"Keterangan"`
}
type DaftarSegmenResponse struct {
	StatusCode int            `json:"StatusCode"`
	StatusDesc string         `json:"StatusDesc"`
	Data       []DaftarSegmen `json:"Data"`
}
