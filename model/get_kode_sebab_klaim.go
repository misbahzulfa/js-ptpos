package model

type DaftarSebabKlaim struct {
	KodeSebabKlaim string `json:"KodeSebabKlaim"`
	NamaSebabKlaim string `json:"NamaSebabKlaim"`
}
type DaftarSebabKlaimResponse struct {
	StatusCode int                `json:"StatusCode"`
	StatusDesc string             `json:"StatusDesc"`
	Data       []DaftarSebabKlaim `json:"Data"`
}
