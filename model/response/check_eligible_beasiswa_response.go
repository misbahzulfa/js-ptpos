package response

type CheckEligibleBeasiswaResponse struct {
	StatusCode            string                  `json:"statusCode"`
	StatusDesc            string                  `json:"statusDesc"`
	CheckEligibleBeasiswa []CheckEligibleBeasiswa `json:"dataCheckEligibleBeasiswa"`
}

type CheckEligibleBeasiswa struct {
	StatusPencarian     string                `json:"statusPencarian"`
	KeteranganPencarian string                `json:"keteranganPencarian"`
	NamaPeserta         string                `json:"namaPeserta"`
	DataDokumenBeasiswa []DataDokumenBeasiswa `json:"dataDokumenBeasiswa"`
	DataProbingBeasiswa []DataProbingBeasiswa `json:"dataProbingBeasiswa"`
}

type DataDokumenBeasiswa struct {
	KodeDokumen string `json:"kodeDokumen"`
	NamaDokumen string `json:"namaDokumen"`
}

type DataProbingBeasiswa struct {
	KodeProbing   string `json:"kodeProbing"`
	NoUrut        string `json:"noUrut"`
	ResponProbing string `json:"responProbing"`
	Kategori      string `json:"kategori"`
}
