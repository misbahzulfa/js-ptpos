package entity

type CheckEligibleBeasiswaEntity struct {
	StatusPencarian     string                             `json:"statusPencarian"`
	KeteranganPencarian string                             `json:"keteranganPencarian"`
	NamaPeserta         string                             `json:"namaPeserta"`
	Data                []CheckEligibleBeasiswaDetilEntity `json:"data"`
	DataProbing         []CheckEligibleDataProbing         `json:"dataProbing"`
}

type CheckEligibleBeasiswaDetilEntity struct {
	KodeDokumen string `json:"kodeDokumen"`
	NamaDokumen string `json:"namaDokumen"`
}

type CheckEligibleDataProbing struct {
	KodeProbing   string `json:"kodeProbing"`
	NoUrut        string `json:"noUrut"`
	ResponProbing string `json:"responProbing"`
}
