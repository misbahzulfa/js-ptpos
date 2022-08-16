package entity

type DaftarJenisBeasiswaEntity struct {
	StatusCode        string              `json:"statusCode"`
	StatusDesc        string              `json:"statusDesc"`
	DataJenisBeasiswa []DataJenisBeasiswa `json:"dataJenisBeasiswa"`
}

type DataJenisBeasiswa struct {
	KodeJenisBeasiswa string `json:"kodeJenisBeasiswa"`
	NamaJenisBeasiswa string `json:"namaJenisBeasiswa"`
}
