package model

type DaftarDokumenSebabKlaim struct {
	KodeDokumen   string `json:"KodeDokumen"`
	NamaDokumen   string `json:"NamaDokumen"`
	FlagMandatory string `json:"FlagMandatory"`
}
type DaftarDokumenSebabKlaimResponse struct {
	StatusCode int                       `json:"StatusCode"`
	StatusDesc string                    `json:"StatusDesc"`
	Data       []DaftarDokumenSebabKlaim `json:"Data"`
}
