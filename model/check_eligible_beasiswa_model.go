package model

type CheckEligibleBeasiswa struct {
	StatusPencarian     string                       `json:"statusPencarian"`
	KeteranganPencarian string                       `json:"keteranganPencarian"`
	Data                []CheckEligibleBeasiswaDetil `json:"Data"`
}

type CheckEligibleBeasiswaDetil struct {
	KodeDokumen string `json:"KodeDokumen"`
	NamaDokumen string `json:"NamaDokumen"`
}
