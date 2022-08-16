package model

type DaftarJenisBeasiswa struct {
	// KodeJenisBeasiswa string `json:"kodeJenisBeasiswa"`
	// NamaJenisBeasiswa string `json:"namaJenisBeasiswa"`

	Ret  string `json:"ret"`
	Data struct {
		KodeJenisBeasiswa string `json:"kodeJenisBeasiswa"`
		NamaJenisBeasiswa string `json:"namaJenisBeasiswa"`
	} `json:"DATA"`
	Msg string `json:"msg"`
}
